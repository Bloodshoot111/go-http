package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)
const(
	defaultMaxIdleConnections = 5
	defaultResponseTimeOut = 5 * time.Second
	defaultConnectionTimeout = 1 * time.Second
)

func (c *httpClient) do(method string, headers http.Header ,url string, body interface{}) (*http.Response, error) {


	fullHeaders :=  c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)

	if err != nil {
		return nil, fmt.Errorf("error while parsing the requestBody: %w", err)
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, fmt.Errorf("error while creating the request: %w", err)
	}

	request.Header = fullHeaders

	client := c.getHttpClient()

	return  client.Do(request)
}

func (c *httpClient) getHttpClient() *http.Client {
	if c.client == nil {
		c.client = &http.Client{
			Timeout: defaultConnectionTimeout * defaultResponseTimeOut,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:  c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	}
	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.disableTimeout {
		return 0
	}
	if c.responseTimeout > 0 {
		return c.responseTimeout
	}
	return defaultResponseTimeOut
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.disableTimeout {
		return 0
	}
	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	}
	return defaultConnectionTimeout
}

func (c *httpClient) getRequestHeaders(headers http.Header) http.Header {
	result := make(http.Header)

	// Add common headers to the request:
	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Add custom headers to the request:
	for header, value := range headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}

	}
	return result
}

func (c* httpClient) getRequestBody(contentType string,body interface{}) ([]byte,error){
	if body == nil {
		return nil,nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}

}
