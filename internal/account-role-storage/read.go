package accountrolestorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Read(ctx context.Context, ops model.ReadAccountRoleOptions, pgOps model.PaginationOptions) ([]model.AccountRole, error) {
	b := s.sq.
		Select(COLUMN_ID, COLUMN_ACCOUNT_ID, COLUMN_ROLE).
		From(TABLE_NAME).
		Where(s.buildWhere(ops))
	if pgOps.Limit != 0 {
		b.Limit(uint64(pgOps.Limit))
	}
	sql, args, err := b.Offset(uint64(pgOps.Offset)).ToSql()
	if err != nil {
		return []model.AccountRole{}, fmt.Errorf("dbaccountrolestorage.Storage.Read: %w", err)
	}
	var rows []model.AccountRole
	if err := s.db.SelectContext(ctx, &rows, sql, args...); err != nil {
		return []model.AccountRole{}, fmt.Errorf("dbaccountrolestorage.Storage.Read: %w", err)
	}
	return rows, nil
}
