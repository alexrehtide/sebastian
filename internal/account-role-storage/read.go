package accountrolestorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Read(ctx context.Context, ops model.ReadAccountRoleOptions, pgOps model.PaginationOptions) (rows []model.AccountRole, err error) {
	b := s.sq.
		Select(COLUMN_ID, COLUMN_ACCOUNT_ID, COLUMN_ROLE).
		From(TABLE_NAME).
		Where(s.buildWhere(ops))
	if pgOps.Limit != 0 {
		b.Limit(uint64(pgOps.Limit))
	}
	sql, args, err := b.Offset(uint64(pgOps.Offset)).ToSql()
	if err != nil {
		return []model.AccountRole{}, fmt.Errorf("accountrolestorage.Storage.Read: %w", err)
	}
	if err := s.getter.DefaultTrOrDB(ctx, s.db).SelectContext(ctx, &rows, sql, args...); err != nil {
		return []model.AccountRole{}, fmt.Errorf("accountrolestorage.Storage.Read: %w", err)
	}
	return rows, nil
}
