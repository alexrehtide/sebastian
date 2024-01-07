package sessionstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Update(ctx context.Context, id uint, ops model.UpdateSessionOptions) error {
	sq := s.sq.Update(TABLE_NAME)
	if ops.AccessToken != "" {
		sq = sq.Set(COLUMN_ACCESS_TOKEN, ops.AccessToken)
	}
	if ops.RefreshToken != "" {
		sq = sq.Set(COLUMN_REFRESH_TOKEN, ops.RefreshToken)
	}
	_, err := sq.
		Set(COLUMN_UPDATED_AT, time.Now()).
		Where(squirrel.Eq{COLUMN_ID: id}).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("sessionstorage.Storage.Update: %w", err)
	}
	return nil
}
