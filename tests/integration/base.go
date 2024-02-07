package integration

import "app/tests/client"

var openApiClient *client.Client

func init() {
	var err error
	openApiClient, err = client.NewClient("http://localhost:8000")
	if err != nil {
		panic(err)
	}
}

func strToPtr(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
