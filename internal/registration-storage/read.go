package registrationstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Read(ctx context.Context, ops model.ReadRegistrationOptions, pgOps model.PaginationOptions) (rows []model.Registration, err error) {
	sql, args, err := s.sq.
		Select(
			COLUMN_ID,
			COLUMN_EMAIL,
			COLUMN_USERNAME,
			COLUMN_PASSWORD,
			COLUMN_VERIFICATION_CODE,
			COLUMN_CREATED_AT,
		).
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		Limit(uint64(pgOps.Limit)).
		Offset(uint64(pgOps.Offset)).
		ToSql()
	if err != nil {
		return []model.Registration{}, fmt.Errorf("registrationstorage.Storage.Read: %w", err)
	}
	if err := s.getter.DefaultTrOrDB(ctx, s.db).SelectContext(ctx, &rows, sql, args...); err != nil {
		return []model.Registration{}, fmt.Errorf("registrationstorage.Storage.Read: %w", err)
	}
	return rows, nil
}
