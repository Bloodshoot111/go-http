package gohttp

import (
	"net/http"
	"sync"
)

type Client interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

type httpClient struct {
	builder *clientBuilder

	client *http.Client
	clientOnce sync.Once
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



