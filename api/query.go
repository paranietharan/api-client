package api

import (
	errors "api-client/error"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Resp interface {
	interface{} // This makes it accept any type, but we'll constrain it in practice
}

type Query[R Resp] struct {
	cli    Client
	Method string
	Path   string
	Body   []byte
}

func NewQuery[R Resp](c Client) Query[R] {
	return Query[R]{
		cli:    c,
		Method: http.MethodGet,
	}
}

func (p Query[R]) WithMethod(method string) Query[R] {
	p.Method = method

	return p
}

func (p Query[R]) WithPath(format string, v ...any) Query[R] {
	p.Path = fmt.Sprintf(format, v...)

	return p
}

func (p Query[R]) WithBody(body any) Query[R] {
	b, _ := json.Marshal(body)
	p.Body = b

	return p
}

func (p Query[R]) Do(ctx context.Context) (R, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent("DoAPIQuery")
	span.SetAttributes(attribute.String("method", p.Method), attribute.String("path", p.Path))

	// p.cli.getLogger().InfoCtx(
	// 	ctx, "start sending request",
	// 	log.String("method", p.Method),
	// 	log.String("path", p.Path),
	// )

	var rsp R
	urlPath, err := url.JoinPath(p.cli.getBaseURL(), p.Path)
	if err != nil {
		return rsp, err
	}

	// p.cli.getLogger().DebugCtx(ctx, "request info",
	// 	log.String("method", p.Method),
	// 	log.String("url", urlPath),
	// 	log.String("request", qkit.B2S(p.Body)),
	// )
	httpRsp, err := p.cli.do(ctx, p.Method, urlPath, bytes.NewReader(p.Body))
	if err != nil {
		fmt.Printf("%v\n", err)
		return rsp, err
	}
	defer func() { _ = httpRsp.Body.Close() }()
	fmt.Println("response")
	fmt.Printf("method %v\n", p.Method)
	fmt.Printf("status %v\n", httpRsp.StatusCode)

	if httpRsp.StatusCode != http.StatusOK && httpRsp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(httpRsp.Body)
		fmt.Printf("requet filed with statuCode: %d\n", httpRsp.StatusCode)
		fmt.Printf("%v\n", body)
	}

	switch httpRsp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusUnauthorized, http.StatusForbidden:
		return rsp, errors.ErrInvalidAuth
	case http.StatusNotFound:
		return rsp, errors.ErrNotFound
	case http.StatusTooManyRequests:
		return rsp, errors.ErrTooManyRequests
	default:
		return rsp, errors.ErrInternalAPICall
	}

	if err := json.NewDecoder(httpRsp.Body).Decode(&rsp); err != nil {
		bb, _ := io.ReadAll(httpRsp.Body)
		err = fmt.Errorf("failed to decode response: %v. body: %v", err, fmt.Sprintf("%v", bb))
		return rsp, err
	}

	return rsp, nil
}
