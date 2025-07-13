package common

import (
	"net/http"

	"resty.dev/v3"
)

type APIClient struct {
	Client   *resty.Client
	BASE_URL string
}

func (a *APIClient) MakeRequest(method string, endpoint string, queryParams map[string]string, result any, payload any, headers map[string]string) (*resty.Response, error) {
	request_url := a.BASE_URL + endpoint
	r := a.Client.R()

	var resp *resty.Response
	var err error

	r.SetQueryParams(queryParams)
	r.SetHeaders(headers)
	r.SetResult(result)

	switch method {
	case http.MethodGet:
		resp, err = r.Get(request_url)
	case http.MethodPut:
		resp, err = r.SetBody(payload).Put(request_url)
	case http.MethodPost:
		resp, err = r.SetBody(payload).Post(request_url)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
