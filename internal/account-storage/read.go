package dbaccountstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Read(ctx context.Context, ops model.ReadAccountOptions, pgOps model.PaginationOptions) ([]model.Account, error) {
	sql, args, err := s.sq.
		Select(COLUMN_ID, COLUMN_EMAIL, COLUMN_PASSWORD_HASH).
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		Limit(uint64(pgOps.Limit)).
		Offset(uint64(pgOps.Offset)).
		ToSql()
	if err != nil {
		return []model.Account{}, fmt.Errorf("dbaccountstorage.Storage.Read: %w", err)
	}
	var rows []model.Account
	if err := s.db.SelectContext(ctx, &rows, sql, args...); err != nil {
		return []model.Account{}, fmt.Errorf("dbaccountstorage.Storage.Read: %w", err)
	}
	return rows, nil
}
