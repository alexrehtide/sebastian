package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func New(ops PostgresOptions) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", ops.User, ops.Password, ops.Host, ops.Port, ops.DBName))
}

type PostgresOptions struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}
