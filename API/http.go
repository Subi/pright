package api

import (
	"io"
	"net/http"
)

type Api struct {
	client http.Client
}

func NewApi() *Api {
	client := http.Client{}
	return &Api{client: client}
}

func (a *Api) Request(method string, url string, headers map[string]string, body io.Reader) (resp *http.Response, err error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return resp, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err = a.client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, err
}
