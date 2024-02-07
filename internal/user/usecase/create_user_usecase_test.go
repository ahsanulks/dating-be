package usecase_test

import (
	"app/infra/encryption"
	"app/internal/adapter/fake"
	"app/internal/user/param/request"
	"app/internal/user/usecase"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	fakeUserDriven := fake.NewFakeUserDriven()
	bcrypt := new(encryption.BcyrpEncryption)
	type args struct {
		ctx   context.Context
		param *request.CreateUser
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "when username less than 5 character it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12zc",
				PhoneNumber: "+628123123123",
				Name:        faker.Name(),
				Password:    generateRandomPassword(12),
				Gender:      "male",
			}},
			wantErr:    true,
			wantErrMsg: "username: must be between 5 and 50 characters in length",
		},
		{
			name: "when username more than 50 character it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12zcd123123hnm-_123jkalsdsajkldjqweqncqawl123hj1312klsahd0085473456324524523424adssa132131qweq",
				PhoneNumber: "+628123123123",
				Name:        faker.Name(),
				Password:    generateRandomPassword(12),
				Gender:      "male",
			}},
			wantErr:    true,
			wantErrMsg: "username: must be between 5 and 50 characters in length",
		},
		{
			name: "when username pattern not valid",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12zcA!",
				PhoneNumber: "+628123123123",
				Name:        faker.Name(),
				Password:    generateRandomPassword(12),
				Gender:      "male",
			}},
			wantErr:    true,
			wantErrMsg: "username: can only contain letter, numbers, underscores, and dashes",
		},
		{
			name: "when gender is not convertable should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+628123123123",
				Name:        faker.Name(),
				Password:    generateRandomPassword(12),
				Gender:      "123",
			}},
			wantErr:    true,
			wantErrMsg: "gender: can only male and female",
		},
		{
			name: "when phoneNumber less than 10 character it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+62123",
				Name:        faker.Name(),
				Password:    generateRandomPassword(12),
				Gender:      "male",
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 15 characters in length",
		},
		{
			name: "when phoneNumber more than 15 character it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+621231231231233333",
				Name:        faker.Name(),
				Password:    generateRandomPassword(12),
				Gender:      "male",
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must be between 10 and 15 characters in length",
		},
		{
			name: "when phoneNumber not have prefix +62 it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "0812311231231",
				Name:        faker.Name(),
				Password:    generateRandomPassword(12),
				Gender:      "male",
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must start with '+62' and only containt number",
		},
		{
			name: "when phoneNumber have prefix +62 but containt other than number, it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+62abc3123123",
				Name:        faker.Name(),
				Password:    generateRandomPassword(12),
				Gender:      "male",
			}},
			wantErr:    true,
			wantErrMsg: "phoneNumber: must start with '+62' and only containt number",
		},
		{
			name: "when Name less than 3 character, it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+628123123123",
				Name:        generateRandomString(2, ""),
				Password:    generateRandomPassword(12),
				Gender:      "male",
			}},
			wantErr:    true,
			wantErrMsg: "name: must be between 3 and 255 characters in length",
		},
		{
			name: "when Name more than 255 character, it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+628123123123",
				Name:        generateRandomString(256, ""),
				Password:    generateRandomPassword(12),
				Gender:      "male",
			}},
			wantErr:    true,
			wantErrMsg: "name: must be between 3 and 255 characters in length",
		},
		{
			name: "when password less than 6 character, it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+628123123123",
				Name:        faker.Name(),
				Password:    "aA2.",
				Gender:      "female",
			}},
			wantErr:    true,
			wantErrMsg: "password: must be between 6 and 64 characters in length",
		},
		{
			name: "when password more than 64 character, it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+628123123123",
				Name:        faker.Name(),
				Password:    generateRandomPassword(65),
				Gender:      "female",
			}},
			wantErr:    true,
			wantErrMsg: "password: must be between 6 and 64 characters in length",
		},
		{
			name: "when password not containt number, it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+628123123123",
				Name:        faker.Name(),
				Password:    "abcAdasd.D",
				Gender:      "female",
			}},
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when password not containt capital char, it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+628123123123",
				Name:        faker.Name(),
				Password:    "abc123.asd",
				Gender:      "female",
			}},
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when password not containt special char, it should return error",
			args: args{context.Background(), &request.CreateUser{
				Username:    "12adsas123_-",
				PhoneNumber: "+628123123123",
				Name:        faker.Name(),
				Password:    "Asd123ASd",
				Gender:      "female",
			}},
			wantErr:    true,
			wantErrMsg: "password: containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when all field not valid, it should return all error message",
			args: args{context.Background(), &request.CreateUser{
				Username:    "",
				PhoneNumber: "",
				Name:        "",
				Password:    "",
				Gender:      "",
			}},
			wantErr:    true,
			wantErrMsg: "username: must be between 5 and 50 characters in length,can only contain letter, numbers, underscores, and dashes;gender: can only male and female;phoneNumber: must be between 10 and 15 characters in length,must start with '+62' and only containt number;name: must be between 3 and 255 characters in length;password: must be between 6 and 64 characters in length,containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) characters",
		},
		{
			name: "when all field valid, it should return id and saved",
			args: args{context.Background(), &request.CreateUser{
				Username:    "Ads123d-s123-_",
				PhoneNumber: "+628123123123",
				Name:        faker.Name(),
				Password:    generateRandomPassword(12),
				Gender:      "MALE",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := usecase.NewUserUsecase(fakeUserDriven, bcrypt)
			gotID, err := uu.CreateUser(tt.args.ctx, tt.args.param)
			assert := assert.New(t)
			if tt.wantErr {
				assert.Error(err)
				assert.True(assertMessagesEqual(tt.wantErrMsg, err.Error()))
			} else {
				gotUser, err := fakeUserDriven.GetByID(tt.args.ctx, gotID)
				assert.NoError(err)
				assert.Equal(gotUser.ID, gotID)
			}
		})
	}
}

func splitAndSortErrorMessage(message string) []string {
	slicedMessages := strings.Split(message, ";")
	for i, msg := range slicedMessages {
		slicedMessages[i] = strings.TrimSpace(msg)
	}

	sort.Strings(slicedMessages)
	return slicedMessages
}

func assertMessagesEqual(message1, message2 string) bool {
	sorted1 := splitAndSortErrorMessage(message1)
	sorted2 := splitAndSortErrorMessage(message2)

	return fmt.Sprintf("%v", sorted1) == fmt.Sprintf("%v", sorted2)
}

func TestCreateUser_withPasswordEncrypted(t *testing.T) {
	fakeUserDriven := fake.NewFakeUserDriven()
	bcrypt := new(encryption.BcyrpEncryption)
	uu := usecase.NewUserUsecase(fakeUserDriven, bcrypt)
	assert := assert.New(t)

	userParam := &request.CreateUser{
		PhoneNumber: "+628123123123",
		Name:        faker.Name(),
		Password:    generateRandomPassword(12),
	}
	gotID, err := uu.CreateUser(context.Background(), userParam)
	assert.NoError(err)

	user, err := fakeUserDriven.GetByID(context.Background(), gotID)
	assert.NoError(err)
	assert.Equal(user.ID, gotID)
	assert.NotEqual(userParam.Password, user.Password)

	err = bcrypt.CompareEncryptedAndData([]byte(user.Password), []byte(userParam.Password))
	assert.NoError(err)
}

func generateRandomString(length int, charset string) string {
	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	charsetLength := big.NewInt(int64(len(charset)))

	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, charsetLength)
		randomString[i] = charset[randomIndex.Int64()]
	}

	return string(randomString)
}

func generateRandomPassword(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>/?~"
	passwordRegex := regexp.MustCompile(`^(.*[A-Z])(.*\d)(.*[^A-Za-z0-9])`)
	var password string
	match := false
	for !match {
		password = generateRandomString(length, charset)
		match = passwordRegex.MatchString(password)
	}

	return password
}
