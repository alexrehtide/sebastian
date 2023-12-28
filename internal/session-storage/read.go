package dbsessionstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Read(ctx context.Context, ops model.ReadSessionOptions, pgOps model.PaginationOptions) ([]model.Session, error) {
	sql, args, err := s.sq.
		Select(COLUMN_ID, COLUMN_ACCOUNT_ID, COLUMN_ACCESS_TOKEN, COLUMN_REFRESH_TOKEN, COLUMN_CREATED_AT, COLUMN_UPDATED_AT).
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		Limit(uint64(pgOps.Limit)).
		Offset(uint64(pgOps.Offset)).
		ToSql()
	if err != nil {
		return []model.Session{}, fmt.Errorf("dbsessionstorage.Storage.Read: %w", err)
	}
	var rows []model.Session
	if err := s.db.SelectContext(ctx, &rows, sql, args...); err != nil {
		return []model.Session{}, fmt.Errorf("dbsessionstorage.Storage.Read: %w", err)
	}
	return rows, nil
}
