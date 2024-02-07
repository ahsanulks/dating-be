-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_tokens (
    id                  BIGSERIAL   PRIMARY KEY,
    user_id             BIGSERIAL   UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    success_login_count INT         DEFAULT 0,
    last_login_at       TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_tokens;
-- +goose StatementEnd
