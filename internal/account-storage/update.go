package accountstorage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Update(ctx context.Context, id uint, ops model.UpdateAccountOptions) error {
	sq := s.sq.Update(TABLE_NAME)
	if ops.Email != "" {
		sq = sq.Set(COLUMN_EMAIL, ops.Email)
	}
	if ops.Username != "" {
		sq = sq.Set(COLUMN_USERNAME, ops.Username)
	}
	if ops.Password != "" {
		sq = sq.Set(COLUMN_PASSWORD, ops.Password)
	}
	_, err := sq.
		Where(squirrel.Eq{COLUMN_ID: id}).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("accountstorage.Storage.Update: %w", err)
	}
	return nil
}
