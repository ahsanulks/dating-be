package database

import (
	"app/internal/user/entity"
	"context"
	"database/sql"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func newMockConn() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestUserRepository_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *entity.User
	}
	tests := []struct {
		name       string
		args       args
		wantId     int64
		wantErr    bool
		expectFunc func(sqlmock.Sqlmock, *entity.User)
	}{
		{
			name: "when given user entity, it should insert Name, phoneNumber, password",
			args: args{
				ctx: context.Background(),
				user: &entity.User{
					ID:          0,
					Name:        faker.Name(),
					Username:    "asdasd123",
					PhoneNumber: faker.Phonenumber(),
					Gender:      entity.Gender("male"),
					Password:    faker.Password(),
				},
			},
			wantId:  123131,
			wantErr: false,
			expectFunc: func(mock sqlmock.Sqlmock, user *entity.User) {
				mock.ExpectQuery("^INSERT INTO users").
					WithArgs(user.Username, user.Name, user.Gender.String(), user.PhoneNumber, user.Password).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(123131))
			},
		},
		{
			name: "when something went wrong in db, it should return error",
			args: args{
				ctx: context.Background(),
				user: &entity.User{
					Name:        faker.Name(),
					PhoneNumber: faker.Phonenumber(),
					Password:    faker.Password(),
				},
			},
			wantId:  0,
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, user *entity.User) {
				mock.ExpectQuery("^INSERT INTO users").
					WithArgs(user.Username, user.Name, user.Gender.String(), user.PhoneNumber, user.Password).
					WillReturnError(errors.New("some database error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, dbMock := newMockConn()
			defer conn.Close()
			pgConn := PostgresDB{
				conn: conn,
			}
			udb := NewUserRepository(&pgConn)

			tt.expectFunc(dbMock, tt.args.user)

			gotId, err := udb.Create(tt.args.ctx, tt.args.user)

			assert := assert.New(t)
			assert.Equal(tt.wantErr, err != nil)
			assert.Equal(tt.wantId, gotId)
			assert.NoError(dbMock.ExpectationsWereMet())
		})
	}
}

func TestUserRepository_UpdateLoginInformation(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *entity.User
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		expectFunc func(sqlmock.Sqlmock, *entity.User)
	}{
		{
			name: "when error on db, it should return error",
			args: args{
				context.Background(),
				&entity.User{
					ID: 123131,
				},
			},
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, user *entity.User) {
				mock.ExpectExec("INSERT INTO user_tokens").
					WithArgs(user.ID).
					WillReturnError(errors.New("some database error"))
			},
		},
		{
			name: "when error on db, it should return error",
			args: args{
				context.Background(),
				&entity.User{
					ID: 3123131321,
				},
			},
			wantErr: false,
			expectFunc: func(mock sqlmock.Sqlmock, user *entity.User) {
				mock.ExpectExec("INSERT INTO user_tokens").
					WithArgs(user.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, dbMock := newMockConn()
			defer conn.Close()
			pgConn := PostgresDB{
				conn: conn,
			}
			udb := NewUserRepository(&pgConn)

			tt.expectFunc(dbMock, tt.args.user)

			err := udb.UpdateLoginInformation(tt.args.ctx, tt.args.user)

			assert := assert.New(t)
			assert.Equal(tt.wantErr, err != nil)
			assert.NoError(dbMock.ExpectationsWereMet())
		})
	}
}

func TestUserRepository_GetByUsername(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name       string
		args       args
		want       *entity.User
		wantErr    bool
		expectFunc func(sqlmock.Sqlmock, *entity.User)
	}{
		{
			name: "when record not found, it should return error",
			args: args{
				context.Background(),
				faker.Username(),
			},
			want:    nil,
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, _ *entity.User) {
				mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{}))
			},
		},
		{
			name: "when error on database, it should return error",
			args: args{
				context.Background(),
				faker.Username(),
			},
			want:    nil,
			wantErr: true,
			expectFunc: func(mock sqlmock.Sqlmock, _ *entity.User) {
				mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("database error"))
			},
		},
		{
			name: "when user found, it should return user",
			args: args{
				context.Background(),
				"testUsername123",
			},
			want: &entity.User{
				ID:          1231321,
				Name:        faker.Name(),
				Username:    "testUsername123",
				PhoneNumber: faker.Phonenumber(),
				Password:    faker.Password(),
				Gender:      entity.Gender("male"),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantErr: false,
			expectFunc: func(mock sqlmock.Sqlmock, expectedUser *entity.User) {
				rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "phone_number", "gender", "created_at", "updated_at"}).
					AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Username, expectedUser.Password, expectedUser.PhoneNumber, "male", expectedUser.CreatedAt, expectedUser.UpdatedAt)

				mock.ExpectQuery("SELECT").WithArgs("testUsername123").WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		conn, dbMock := newMockConn()
		defer conn.Close()
		pgConn := PostgresDB{
			conn: conn,
		}
		udb := NewUserRepository(&pgConn)

		tt.expectFunc(dbMock, tt.want)

		got, err := udb.GetByUsername(tt.args.ctx, tt.args.username)

		assert := assert.New(t)
		assert.Equal(tt.wantErr, err != nil)
		assert.Equal(tt.want, got)
		assert.NoError(dbMock.ExpectationsWereMet())
	}
}
