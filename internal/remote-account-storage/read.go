package remoteaccountstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Read(ctx context.Context, ops model.ReadRemoteAccountOptions, pgOps model.PaginationOptions) (rows []model.RemoteAccount, err error) {
	sql, args, err := s.sq.
		Select(COLUMN_ID, COLUMN_ACCOUNT_ID, COLUMN_REMOTE_ID, COLUMN_REMOTE_EMAIL, COLUMN_PLATFORM).
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		Limit(uint64(pgOps.Limit)).
		Offset(uint64(pgOps.Offset)).
		ToSql()
	if err != nil {
		return []model.RemoteAccount{}, fmt.Errorf("remoteaccountstorage.Storage.Read: %w", err)
	}
	if err := s.getter.DefaultTrOrDB(ctx, s.db).SelectContext(ctx, &rows, sql, args...); err != nil {
		return []model.RemoteAccount{}, fmt.Errorf("remoteaccountstorage.Storage.Read: %w", err)
	}
	return rows, nil
}
