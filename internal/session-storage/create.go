package dbsessionstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Create(ctx context.Context, ops model.CreateSessionOptions) (id uint, err error) {
	err = s.sq.
		Insert(TABLE_NAME).
		Columns(COLUMN_ACCOUNT_ID, COLUMN_ACCESS_TOKEN, COLUMN_REFRESH_TOKEN, COLUMN_CREATED_AT, COLUMN_UPDATED_AT).
		Values(ops.AccountID, ops.AccessToken, ops.RefreshToken, time.Now(), time.Now()).
		Suffix(fmt.Sprintf("RETURNING %s", COLUMN_ID)).
		ScanContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("dbsessionstorage.Storage.Create: %w", err)
	}
	return id, nil
}
