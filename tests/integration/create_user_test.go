package integration

import (
	"app/tests/client"
	"context"
	"io"
	"math/rand"
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	username := generateUsername(10)
	password := generatePassword(20)
	type arg struct {
		ctx  context.Context
		body client.UserCreateUserJSONRequestBody
	}
	tests := []struct {
		name         string
		arg          arg
		wantHttpCode int
		wantMessage  string
	}{
		{
			name: "when input valid, it should return 200",
			arg: arg{
				ctx: context.Background(),
				body: client.UserCreateUserJSONRequestBody{
					Gender:      strToPtr("male"),
					Name:        strToPtr(faker.Name()),
					Password:    strToPtr(password),
					PhoneNumber: strToPtr(generatePhoneNumber()),
					Username:    strToPtr(username),
				},
			},
			wantHttpCode: http.StatusOK,
			wantMessage:  "id",
		},
		{
			name: "when input valid but previous username already registered, it should return 409",
			arg: arg{
				ctx: context.Background(),
				body: client.UserCreateUserJSONRequestBody{
					Gender:      strToPtr("male"),
					Name:        strToPtr(faker.Name()),
					Password:    strToPtr(password),
					PhoneNumber: strToPtr(generatePhoneNumber()),
					Username:    strToPtr(username),
				},
			},
			wantHttpCode: http.StatusConflict,
			wantMessage:  "DuplicateResource",
		},
		{
			name: "when phone not valid, it should return 400",
			arg: arg{
				ctx: context.Background(),
				body: client.UserCreateUserJSONRequestBody{
					Gender:      strToPtr("male"),
					Name:        strToPtr(faker.Name()),
					Password:    strToPtr(password),
					PhoneNumber: strToPtr(faker.Phonenumber()),
					Username:    strToPtr(generateUsername(10)),
				},
			},
			wantHttpCode: http.StatusBadRequest,
			wantMessage:  "phoneNumber",
		},
		{
			name: "when password not valid, it should return 400",
			arg: arg{
				ctx: context.Background(),
				body: client.UserCreateUserJSONRequestBody{
					Gender:      strToPtr("male"),
					Name:        strToPtr(faker.Name()),
					Password:    strToPtr(faker.Password()),
					PhoneNumber: strToPtr(generatePhoneNumber()),
					Username:    strToPtr(generateUsername(10)),
				},
			},
			wantHttpCode: http.StatusBadRequest,
			wantMessage:  "password",
		},
		{
			name: "when username not valid, it should return 400",
			arg: arg{
				ctx: context.Background(),
				body: client.UserCreateUserJSONRequestBody{
					Gender:      strToPtr("male"),
					Name:        strToPtr(faker.Name()),
					Password:    strToPtr(generatePassword(15)),
					PhoneNumber: strToPtr(generatePhoneNumber()),
					Username:    strToPtr("asd123@13#"),
				},
			},
			wantHttpCode: http.StatusBadRequest,
			wantMessage:  "username",
		},
		{
			name: "when gender not valid, it should return 400",
			arg: arg{
				ctx: context.Background(),
				body: client.UserCreateUserJSONRequestBody{
					Gender:      strToPtr("asdqwe"),
					Name:        strToPtr(faker.Name()),
					Password:    strToPtr(generatePassword(15)),
					PhoneNumber: strToPtr(generatePhoneNumber()),
					Username:    strToPtr(generateUsername(12)),
				},
			},
			wantHttpCode: http.StatusBadRequest,
			wantMessage:  "gender",
		},
		{
			name: "when name not valid, it should return 400",
			arg: arg{
				ctx: context.Background(),
				body: client.UserCreateUserJSONRequestBody{
					Gender:      strToPtr("asdqwe"),
					Name:        strToPtr("as"),
					Password:    strToPtr(generatePassword(15)),
					PhoneNumber: strToPtr(generatePhoneNumber()),
					Username:    strToPtr(generateUsername(12)),
				},
			},
			wantHttpCode: http.StatusBadRequest,
			wantMessage:  "name",
		},
	}
	for _, tt := range tests {
		resp, err := openApiClient.UserCreateUser(tt.arg.ctx, tt.arg.body)
		body, _ := io.ReadAll(resp.Body)

		assert := assert.New(t)
		assert.NoError(err)
		assert.Equal(tt.wantHttpCode, resp.StatusCode)
		assert.Contains(string(body), tt.wantMessage)
	}

}

func generateUsername(length int) string {
	rand.New(rand.NewSource(10000))

	// Define the character set for the username
	charSet := "abcdefghijklmnopqrstuvwxyz0123456789_-"

	// Generate the username by randomly selecting characters from the character set
	username := make([]byte, length)
	for i := range username {
		username[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(username)
}

func generatePassword(length int) string {
	rand.New(rand.NewSource(10000))

	// Define character sets for each type of character
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specialChars := "!@#$%^&*()-_=+[]{};:'\",.<>?/|\\`~"

	// Combine all character sets
	charSet := lowercase + uppercase + numbers + specialChars

	// Initialize slices to track the types of characters used in the password
	var password []byte
	hasLower, hasUpper, hasNumber, hasSpecial := false, false, false, false

	// Generate the password
	for len(password) < length {
		char := charSet[rand.Intn(len(charSet))]

		// Ensure each character type requirement is met
		if !hasLower && charSetContains(char, lowercase) {
			hasLower = true
		} else if !hasUpper && charSetContains(char, uppercase) {
			hasUpper = true
		} else if !hasNumber && charSetContains(char, numbers) {
			hasNumber = true
		} else if !hasSpecial && charSetContains(char, specialChars) {
			hasSpecial = true
		}

		// Add the character to the password
		password = append(password, char)
	}

	// Ensure all character type requirements are met
	if !hasLower || !hasUpper || !hasNumber || !hasSpecial {
		// Regenerate password if any requirement is not met
		return generatePassword(length)
	}

	// Shuffle the password characters
	shuffle(password)

	return string(password)
}

func shuffle(slice []byte) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// charSetContains checks if the given character exists in the character set.
func charSetContains(char byte, charSet string) bool {
	for i := range charSet {
		if charSet[i] == char {
			return true
		}
	}
	return false
}

func generatePhoneNumber() string {
	rand.New(rand.NewSource(10000))

	// Define the prefix and length of the phone number excluding the prefix
	prefix := "+62"
	numberLength := 9

	// Generate the digits for the phone number
	var phoneNumberDigits []byte
	for i := 0; i < numberLength; i++ {
		digit := byte(rand.Intn(10)) + '0' // Random digit from '0' to '9'
		phoneNumberDigits = append(phoneNumberDigits, digit)
	}

	// Combine the prefix and digits to form the complete phone number
	phoneNumber := prefix + string(phoneNumberDigits)

	return phoneNumber
}
