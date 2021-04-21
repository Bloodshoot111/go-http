package examples

import (
	"fmt"
)

type Endpoints struct {
	CurrentUserUrl string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error){
	response, err := httpClient.Get("https://api.github.com",nil)
	if err != nil {
		return nil, fmt.Errorf("unable to send request to github.com %w", err)
	}

	fmt.Println(fmt.Sprintf("Status Code: %d", response.StatusCode()))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status()))
	fmt.Println(fmt.Sprintf("Body: %s", response.String()))

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil{
		return nil, fmt.Errorf("unable to unmarshall responseBody %w", err)
	}

	fmt.Println(fmt.Sprintf("Repository Url: %s", endpoints.RepositoryUrl))
	fmt.Println(fmt.Sprintf("Authorizations Url: %s", endpoints.AuthorizationsUrl))
	fmt.Println(fmt.Sprintf("Current User Url: %s", endpoints.CurrentUserUrl))
	return &endpoints, nil
}
