package database

import (
	"app/internal/user/entity"
	"app/internal/user/port/driven"
	"context"
	"database/sql"
)

type UserRepository struct {
	db *PostgresDB
}

var (
	_ driven.UserWriter = new(UserRepository)
	_ driven.UserGetter = new(UserRepository)
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

// UpdateLoginInformation implements driven.UserWriter.
func (ur *UserRepository) UpdateLoginInformation(ctx context.Context, user *entity.User) error {
	_, err := ur.db.Conn().ExecContext(ctx, `
		INSERT INTO
			user_tokens (user_id, success_login_count, last_login_at)
		VALUES
    		($1, 1, NOW())
		ON CONFLICT (user_id)
		DO UPDATE SET
    		success_login_count = user_tokens.success_login_count + 1,
    		last_login_at = NOW()`, user.ID)
	return err
}

// GetByUsername implements driven.UserGetter.
func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	rows, err := ur.db.Conn().QueryContext(ctx, `
		SELECT
			id,
			name,
			username,
			password,
			phone_number,
			gender,
			created_at,
			updated_at
		FROM
			users
		WHERE
			username = $1
		LIMIT
			1
	`, username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var user entity.User
	if rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Password,
			&user.PhoneNumber,
			&user.Gender,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	} else {
		return nil, sql.ErrNoRows
	}

	return &user, err
}
