package accountrolestorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
		sq: squirrel.StatementBuilder.RunWith(db).PlaceholderFormat(squirrel.Dollar),
	}
}

type Storage struct {
	db *sqlx.DB
	sq squirrel.StatementBuilderType
}
