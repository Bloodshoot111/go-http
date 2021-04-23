package gohttp

import (
	"fmt"
	"testing"
)


func TestGetRequestBody(t *testing.T){

	client := httpClient{}
	t.Run("NoBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)

		// Validation
		if err != nil {
			t.Error("no error expected when passing a nil body")
		}

		if body != nil {
			t.Error("no body expected when passing a nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string {"one","two"}
		body, err := client.getRequestBody("application/json",requestBody)

		if err != nil {
			t.Error("no error expected when passing a valid body")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		requestBody := []string {"one","two"}
		body, err := client.getRequestBody("application/xml",requestBody)

		if err != nil {
			t.Error("no error expected when passing a valid body")
		}

		if string(body) != `<string>one</string><string>two</string>` {
			t.Error("invalid xml body obtained")
		}
	})

	t.Run("BodyWithOctetStream", func(t *testing.T) {
		requestBody := []string {"testFile.txt"}
		body, err := client.getRequestBody("application/octet-stream",requestBody)

		if err != nil {
			t.Error("no error expected when passing a valid body")
		}
		fmt.Println(string(body))
		if string(body) != `This is a testFile` {
			t.Error("invalid octet-stream body obtained")
		}
	})

	t.Run("BodyWithJsonasDefault", func(t *testing.T) {
		requestBody := []string {"one","two"}
		body, err := client.getRequestBody("",requestBody)

		if err != nil {
			t.Error("no error expected when passing a valid body")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})
}