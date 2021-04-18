package gohttp

import "net/http"

type Client interface {
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

type httpClient struct {
	Headers http.Header
}

func(c* httpClient) SetHeaders(headers http.Header){
	c.Headers = headers
}

func New() Client {
	client := &httpClient{}
	return client
}

func (c* httpClient) Get(url string, headers http.Header) (*http.Response, error){
	return c.do(http.MethodGet, headers, url, nil)
}

func (c* httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodPost, headers, url, body)
}

func (c* httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodPut, headers, url, body)
}

func (c* httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error){
	return c.do(http.MethodPatch, headers, url, body)
}

func (c* httpClient) Delete(url string, headers http.Header) (*http.Response, error){
	return c.do(http.MethodDelete, headers, url, nil)
}



