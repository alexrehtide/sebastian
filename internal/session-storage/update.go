package sessionstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Update(ctx context.Context, id uint, ops model.UpdateSessionOptions) error {
	_, err := s.sq.
		Update(TABLE_NAME).
		Set(COLUMN_ACCESS_TOKEN, ops.AccessToken).
		Set(COLUMN_REFRESH_TOKEN, ops.RefreshToken).
		Set(COLUMN_UPDATED_AT, time.Now()).
		Where(squirrel.Eq{COLUMN_ID: id}).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("sessionstorage.Storage.Update: %w", err)
	}
	return nil
}
