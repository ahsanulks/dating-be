package database

import (
	"app/configs"
	"database/sql"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	conn *sql.DB
}

func NewPostgresDB(c *configs.DBConfig, logger log.Logger) (*PostgresDB, func()) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Hostname, c.Port, c.DB)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("cannot connect to db")
	}
	if err := conn.Ping(); err != nil {
		panic("cannot ping db")
	}
	_ = logger.Log(log.LevelInfo, "msg", "connecting to db")
	cleanup := func() {
		_ = logger.Log(log.LevelInfo, "msg", "closing db connection")
		_ = conn.Close()
	}
	return &PostgresDB{
		conn: conn,
	}, cleanup
}

func (db *PostgresDB) Conn() *sql.DB {
	return db.conn
}
