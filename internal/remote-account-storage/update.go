package remoteaccountstorage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Update(ctx context.Context, id uint, ops model.UpdateRemoteAccountOptions) error {
	_, err := s.sq.
		Update(TABLE_NAME).
		Set(COLUMN_ACCOUNT_ID, ops.AccountID).
		Where(squirrel.Eq{COLUMN_ID: id}).
		RunWith(s.getter.DefaultTrOrDB(ctx, s.db)).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("remoteaccountstorage.Storage.Update: %w", err)
	}
	return nil
}
