package integration

import (
	"app/tests/client"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestGenerateUserToken(t *testing.T) {
	loginParams := client.UserCreateUserTokenJSONRequestBody{
		Password: strToPtr(generateUsername(10)),
		Username: strToPtr(generateUsername(12)),
	}

	loginResponse, err := openApiClient.UserCreateUserToken(context.Background(), loginParams)
	assert := assert.New(t)

	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, loginResponse.StatusCode)

	body, _ := io.ReadAll(loginResponse.Body)
	assert.Contains(string(body), "username")

	username := generateUsername(10)
	password := generatePassword(20)
	resp, err := openApiClient.UserCreateUser(context.Background(), client.UserCreateUserJSONRequestBody{
		Gender:      strToPtr("male"),
		Name:        strToPtr(faker.Name()),
		Password:    strToPtr(password),
		PhoneNumber: strToPtr(generatePhoneNumber()),
		Username:    strToPtr(username),
	})

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	loginParams = client.UserCreateUserTokenJSONRequestBody{
		Password: strToPtr(password),
		Username: strToPtr(username),
	}
	loginResponse, err = openApiClient.UserCreateUserToken(context.Background(), loginParams)

	assert.NoError(err)
	assert.Equal(http.StatusOK, loginResponse.StatusCode)
	body, _ = io.ReadAll(loginResponse.Body)
	assert.Contains(string(body), "token")
	assert.Contains(string(body), "expiresIn")
	assert.Contains(string(body), "Bearer")
}
