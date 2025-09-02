package api

import (
	"api-client/config"
	"net/http"
)

type Response = interface{}

type APIClient struct {
	Client *RestClient
}

func NewAPIClient(cfg config.Config) *APIClient {
	return &APIClient{
		Client: NewRestClient(RestClientConfig{
			BaseUrl: cfg.BaseUrl,
			Dump:    cfg.Dump,
		}),
	}
}

func (c *APIClient) Post(body interface{}) Query[Response] {
	return NewQuery[Response](c.Client).
		WithMethod(http.MethodPost).
		WithPath("/orders").
		WithBody(body)
}

func (c *APIClient) Put(body interface{}) Query[Response] {
	return NewQuery[Response](c.Client).
		WithMethod(http.MethodPost).
		WithPath("/orders").
		WithBody(body)
}

func (c *APIClient) Patch(body interface{}) Query[Response] {
	return NewQuery[Response](c.Client).
		WithMethod(http.MethodPost).
		WithPath("/orders").
		WithBody(body)
}

func (c *APIClient) Delete(body interface{}) Query[Response] {
	return NewQuery[Response](c.Client).
		WithMethod(http.MethodPost).
		WithPath("/orders").
		WithBody(body)
}

func (c *APIClient) Get(body interface{}) Query[Response] {
	return NewQuery[Response](c.Client).
		WithMethod(http.MethodPost).
		WithPath("/orders").
		WithBody(body)
}
