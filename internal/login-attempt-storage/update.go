package loginattemptstorage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Update(ctx context.Context, id uint, ops model.UpdateLoginAttemptOptions) error {
	_, err := s.sq.
		Update(TABLE_NAME).
		Set(COLUMN_COUNT, ops.Count).
		Set(COLUMN_LAST_FAILED, ops.LastFailed).
		Where(squirrel.Eq{COLUMN_ID: id}).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("loginattemptstorage.Storage.Update: %w", err)
	}
	return nil
}
