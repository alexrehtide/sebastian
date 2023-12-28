package dbaccountrolestorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Read(ctx context.Context, ops model.ReadAccountRoleOptions, pgOps model.PaginationOptions) (rows []model.AccountRole, err error) {
	sql, args, err := s.sq.
		Select(COLUMN_ID, COLUMN_ACCOUNT_ID, COLUMN_ROLE).
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		Limit(uint64(pgOps.Limit)).
		Offset(uint64(pgOps.Offset)).
		ToSql()
	if err != nil {
		return []model.AccountRole{}, fmt.Errorf("dbaccountrolestorage.Storage.Read: %w", err)
	}
	if err := s.db.SelectContext(ctx, rows, sql, args...); err != nil {
		return []model.AccountRole{}, fmt.Errorf("dbaccountrolestorage.Storage.Read: %w", err)
	}
	return
}
