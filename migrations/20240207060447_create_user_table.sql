-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id              BIGSERIAL    PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    username        VARCHAR(50) NOT NULL UNIQUE NOT NULL,
    password        VARCHAR(255) NOT NULL,
    phone_number    VARCHAR(15)  UNIQUE NOT NULL,
    gender          VARCHAR(10)  NOT NULL,     
    created_at      TIMESTAMPTZ  DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
