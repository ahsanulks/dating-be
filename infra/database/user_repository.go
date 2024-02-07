package database

import (
	"app/internal/user/entity"
	"app/internal/user/port/driven"
	"context"
)

type UserRepository struct {
	db *PostgresDB
}

var (
	_ driven.UserWriter = new(UserRepository)
)

func NewUserRepository(db *PostgresDB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create implements driven.UserWriter.
func (ur *UserRepository) Create(ctx context.Context, user *entity.User) (id int64, err error) {
	err = ur.db.Conn().QueryRowContext(ctx, `
	INSERT INTO
		users (username, name, gender, phone_number, password)
	VALUES
		($1, $2, $3, $4, $5)
	RETURNING
		id
	`, user.Username, user.Name, user.Gender.String(), user.PhoneNumber, user.Password).Scan(&id)
	return
}
