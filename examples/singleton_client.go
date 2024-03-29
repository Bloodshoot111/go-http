package examples

import (
	"github.com/bloodshoot111/go-http/gohttp"
	"github.com/bloodshoot111/go-http/gomime"
	"net/http"
	"time"
)

var(
	httpClient = getHttpClient()
)


func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Fedes-Computer").
		Build()
	return client
}