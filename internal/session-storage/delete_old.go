package sessionstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
)

func (s *Storage) DeleteOld(ctx context.Context, updatedAtLt time.Time) error {
	_, err := s.sq.
		Delete(TABLE_NAME).
		Where(squirrel.Lt{COLUMN_UPDATED_AT: updatedAtLt}).
		QueryContext(ctx)
	if err != nil {
		return fmt.Errorf("sessionstorage.Storage.DeleteOld: %w", err)
	}
	return nil
}
