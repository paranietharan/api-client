package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

type Client interface {
	getBaseURL() string
	do(ctx context.Context, method, url string, body io.Reader) (*http.Response, error)
}

type RestClientConfig struct {
	BaseUrl string
	Dump    bool
}

type RestClient struct {
	RestClientConfig
	httpC http.Client
}

func NewRestClient(config RestClientConfig) *RestClient {
	return &RestClient{
		RestClientConfig: config,
		httpC: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

func (c *RestClient) newAuthorizedRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (c *RestClient) do(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	req, err := c.newAuthorizedRequest(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	if c.RestClientConfig.Dump {
		bbReqDump, _ := httputil.DumpRequestOut(req, true)
		fmt.Printf("request ready to be sent")
		fmt.Printf("method %s\n", req.Method)
		fmt.Printf("url %s\n", req.URL.String())
		fmt.Printf("request %v\n", bbReqDump)
	}

	rsp, err := c.httpC.Do(req)
	if err != nil {
		return nil, err
	}

	if c.RestClientConfig.Dump {
		//bbRspDump, _ := httputil.DumpResponse(rsp, true)
	}

	return rsp, nil
}

func (c *RestClient) getBaseURL() string {
	return c.RestClientConfig.BaseUrl
}
