package model

import "time"

type Log struct {
	ID        uint      `db:"id"`
	AccountID uint      `db:"account_id"`
	SessionID uint      `db:"session_id"`
	Message   string    `db:"message"`
	Data      string    `db:"data"`
	Level     string    `db:"level"`
	CreatedAt time.Time `db:"created_at"`
}

type CreateLogOptions struct {
	AccountID uint      `db:"account_id"`
	SessionID uint      `db:"session_id"`
	Message   string    `db:"message"`
	Data      string    `db:"data"`
	Level     string    `db:"level"`
	CreatedAt time.Time `db:"created_at"`
}
