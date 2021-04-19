package gohttp

import (
	"net/http"
	"time"
)

type Client interface {
	SetHeaders(headers http.Header)
	SetConnectionTimeout(timeout time.Duration)
	SetResponseTimeout(timeout time.Duration)
	SetMaxIdleConnections(maxIdleConnections int)
	DisableTimeouts(disable bool)
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

type httpClient struct {
	client *http.Client

	maxIdleConnections int
	connectionTimeout time.Duration
	responseTimeout time.Duration
	disableTimeout bool

	Headers http.Header
}

func(c* httpClient) SetHeaders(headers http.Header){
	c.Headers = headers
}

func(c* httpClient) SetConnectionTimeout(timeout time.Duration){
	c.connectionTimeout = timeout
}

func(c* httpClient) SetResponseTimeout(timeout time.Duration){
	c.responseTimeout = timeout
}

func(c* httpClient) SetMaxIdleConnections(maxIdleConnections int){
	c.maxIdleConnections = maxIdleConnections
}

func(c* httpClient) DisableTimeouts(disable bool){
	c.disableTimeout = disable
}

func New() Client {
	httpClient := &httpClient{}
	return httpClient
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



