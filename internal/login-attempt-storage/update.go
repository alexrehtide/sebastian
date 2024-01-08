package loginattemptstorage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Update(ctx context.Context, id uint, ops model.UpdateLoginAttemptOptions) error {
	sq := s.sq.Update(TABLE_NAME)
	if ops.Count != 0 {
		sq = sq.Set(COLUMN_COUNT, ops.Count)
	}
	if !ops.LastFailed.IsZero() {
		sq = sq.Set(COLUMN_LAST_FAILED, ops.LastFailed)
	}
	_, err := sq.
		Where(squirrel.Eq{COLUMN_ID: id}).
		RunWith(s.getter.DefaultTrOrDB(ctx, s.db)).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("loginattemptstorage.Storage.Update: %w", err)
	}
	return nil
}
