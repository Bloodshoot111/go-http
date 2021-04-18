package gohttp

type Client interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

type httpClient struct {}

func New() Client {
	client := &httpClient{}
	return client
}

func (c* httpClient) Get(){}

func (c* httpClient) Post(){}

func (c* httpClient) Put(){}

func (c* httpClient) Patch(){}

func (c* httpClient) Delete(){}



