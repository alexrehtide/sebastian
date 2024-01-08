package accountstorage

import (
	"github.com/Masterminds/squirrel"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB, c *trmsqlx.CtxGetter) *Storage {
	return &Storage{
		db:     db,
		sq:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		getter: c,
	}
}

type Storage struct {
	db     *sqlx.DB
	sq     squirrel.StatementBuilderType
	getter *trmsqlx.CtxGetter
}
