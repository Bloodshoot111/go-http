package main

import (
	"fmt"
	"github.com/bloodshoot111/go-http/gohttp"
	"io/ioutil"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func main() {
	getTest()
	postTest(User{
		FirstName: "Test",
		LastName:  "User",
	})
}

func getTest() {

	client := gohttp.New()
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func postTest(user User) {

	client := gohttp.New()
	response, err := client.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

