package infra

import (
	"app/configs"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	conn *pgxpool.Pool
}

func NewPostgresDB(c *configs.DBConfig, logger log.Logger) (*PostgresDB, func()) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", c.User, c.Password, c.Hostname, c.Port, c.DB)
	conn, err := pgxpool.New(
		context.Background(),
		dsn,
	)
	if err != nil {
		panic("cannot connect to db")
	}
	if err := conn.Ping(context.TODO()); err != nil {
		panic("cannot ping db")
	}
	logger.Log(log.LevelInfo, "msg", "connecting to db")
	cleanup := func() {
		logger.Log(log.LevelInfo, "msg", "closing db connection")
		conn.Close()
	}
	return &PostgresDB{
		conn: conn,
	}, cleanup
}

func (db *PostgresDB) Conn() *pgxpool.Pool {
	return db.conn
}
