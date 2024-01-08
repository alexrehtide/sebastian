package loginattemptstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Read(ctx context.Context, ops model.ReadLoginAttemptOptions, pgOps model.PaginationOptions) (rows []model.LoginAttempt, err error) {
	sql, args, err := s.sq.
		Select(COLUMN_ID, COLUMN_IP, COLUMN_COUNT, COLUMN_LAST_FAILED).
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		Limit(uint64(pgOps.Limit)).
		Offset(uint64(pgOps.Offset)).
		ToSql()
	if err != nil {
		return []model.LoginAttempt{}, fmt.Errorf("loginattemptstorage.Storage.Read: %w", err)
	}
	if err := s.getter.DefaultTrOrDB(ctx, s.db).SelectContext(ctx, &rows, sql, args...); err != nil {
		return []model.LoginAttempt{}, fmt.Errorf("loginattemptstorage.Storage.Read: %w", err)
	}
	return rows, nil
}
