package usecase_test

import (
	"app/infra/encryption"
	"app/internal/adapter/fake"
	"app/internal/user/entity"
	"app/internal/user/param/request"
	"app/internal/user/param/response"
	"app/internal/user/usecase"
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserWriterUsecase_GenerateUsertoken(t *testing.T) {
	fakeUserDriven := fake.NewFakeUserDriven()
	validPassword := faker.Password()
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(validPassword), bcrypt.DefaultCost)
	user := &entity.User{
		Username: faker.Username(),
		Name:     faker.Name(),
		Password: string(encryptedPassword),
	}

	assert := assert.New(t)
	_, err := fakeUserDriven.Create(context.Background(), user)
	assert.NoError(err)

	invalidUser := &entity.User{
		Username: "wrongUsername",
		Name:     faker.Name(),
		Password: string(encryptedPassword),
	}
	_, err = fakeUserDriven.Create(context.Background(), invalidUser)
	assert.NoError(err)

	type args struct {
		ctx    context.Context
		params *request.GenerateUserToken
	}
	tests := []struct {
		name    string
		args    args
		want    *response.Token
		wantErr bool
	}{
		{
			name: "when user not found, it should return error",
			args: args{
				context.Background(),
				&request.GenerateUserToken{
					Username: faker.Username(),
					Password: faker.Password(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when user password error, it should return error",
			args: args{
				context.Background(),
				&request.GenerateUserToken{
					Username: user.Username,
					Password: faker.Password(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when generate token error, it should return error",
			args: args{
				context.Background(),
				&request.GenerateUserToken{
					Username: invalidUser.Username,
					Password: validPassword,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when error record last login count, it should return token",
			args: args{
				context.WithValue(context.Background(), fake.ContextType("token_error"), true),
				&request.GenerateUserToken{
					Username: user.Username,
					Password: validPassword,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success, it should return token",
			args: args{
				context.Background(),
				&request.GenerateUserToken{
					Username: user.Username,
					Password: validPassword,
				},
			},
			want: &response.Token{
				Token:     "1231313213213131",
				ExpiresIn: 3600,
				Type:      "Bearer",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := usecase.NewUserWriterUsecase(fakeUserDriven, new(encryption.BcryptEncryption), fakeUserDriven, new(fake.FakeTokenProvider))
			result, err := uu.GenerateUserToken(tt.args.ctx, tt.args.params)

			assert.Equal(tt.wantErr, err != nil)
			assert.Equal(tt.want, result)
		})
	}
}
