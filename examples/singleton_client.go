package examples

import (
	"github.com/bloodshoot111/go-http/gohttp"
	"time"
)

var(
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3*time.Second).
		Build()
	return client

}
