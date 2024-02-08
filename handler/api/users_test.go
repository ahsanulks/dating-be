package api

import (
	v1 "app/api/v1"
	"app/internal/user/port/driver"
	"app/tests/fake"
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
)

func TestUserApiHandler_CreateUser(t *testing.T) {
	type fields struct {
		userWriter driver.UserWriterUsecase
		log        log.Logger
	}
	type args struct {
		ctx    context.Context
		params *v1.CreateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "when create user error, it should return error",
			fields: fields{
				userWriter: new(fake.FakeUserUsecase),
				log:        log.DefaultLogger,
			},
			args: args{
				ctx: context.Background(),
				params: &v1.CreateUserRequest{
					Username:    "test123",
					Password:    faker.Password(),
					Name:        faker.Name(),
					Gender:      faker.Gender(),
					PhoneNumber: faker.Phonenumber(),
				},
			},
			wantErr: true,
		},
		{
			name: "when create user success, it should return user id",
			fields: fields{
				userWriter: new(fake.FakeUserUsecase),
				log:        log.DefaultLogger,
			},
			args: args{
				ctx: context.Background(),
				params: &v1.CreateUserRequest{
					Username:    faker.Username(),
					Password:    faker.Password(),
					Name:        faker.Name(),
					Gender:      faker.Gender(),
					PhoneNumber: faker.Phonenumber(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewUserApiHandler(tt.fields.userWriter, tt.fields.log)
			got, err := h.CreateUser(tt.args.ctx, tt.args.params)
			assert := assert.New(t)
			if tt.wantErr {
				assert.Error(err)
				assert.Empty(got)
			} else {
				assert.NotEmpty(t, got.Id)
			}
		})
	}
}

func TestUserApiHandler_CreateUserToken(t *testing.T) {
	type fields struct {
		UnimplementedUserServer v1.UnimplementedUserServer
		userWriter              driver.UserWriterUsecase
		log                     log.Logger
	}
	type args struct {
		ctx    context.Context
		params *v1.CreateUserTokenRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "when generate token error, it should return error",
			fields: fields{
				userWriter: new(fake.FakeUserUsecase),
				log:        log.DefaultLogger,
			},
			args: args{
				ctx: context.Background(),
				params: &v1.CreateUserTokenRequest{
					Username: "test123",
					Password: faker.Password(),
				},
			},
			wantErr: true,
		},
		{
			name: "when generate token success, it should return token",
			fields: fields{
				userWriter: new(fake.FakeUserUsecase),
				log:        log.DefaultLogger,
			},
			args: args{
				ctx: context.Background(),
				params: &v1.CreateUserTokenRequest{
					Username: faker.Username(),
					Password: faker.Password(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewUserApiHandler(tt.fields.userWriter, tt.fields.log)
			got, err := h.CreateUserToken(tt.args.ctx, tt.args.params)
			assert := assert.New(t)
			if tt.wantErr {
				assert.Error(err)
				assert.Nil(got)
			} else {
				assert.NotEmpty(got)
				assert.NotEmpty(got.Token)
				assert.NotEmpty(got.ExpiresIn)
				assert.Equal("Bearer", got.Type)
			}
		})
	}
}
