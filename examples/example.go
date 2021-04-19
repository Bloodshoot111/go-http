package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/bloodshoot111/go-http/gohttp"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	getTest()
	postTest(User{
		FirstName: "Test",
		LastName:  "User",
	})
}

func getTest() {

	client := gohttp.NewBuilder().
		DisableTimeouts(false).
		SetConnectionTimeout(10 * time.Second).
		SetMaxIdleConnections(10).
		SetResponseTimeout(5 * time.Second).
		Build()
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func postTest(user User) {

	client := gohttp.NewBuilder().
		DisableTimeouts(false).
		SetConnectionTimeout(10 * time.Second).
		SetMaxIdleConnections(10).
		SetResponseTimeout(5 * time.Second).
		Build()
	response, err := client.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
