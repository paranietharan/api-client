package api

import (
	"api-client/config"
	"api-client/dto"
	"net/http"
)

type Response = dto.Order

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

func (c *APIClient) SubmitOrder(body dto.Order) Query[Response] {
	return NewQuery[Response](c.Client).
		WithMethod(http.MethodPost).
		WithPath("/orders").
		WithBody(body)
}
