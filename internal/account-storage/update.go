package accountstorage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Update(ctx context.Context, id uint, ops model.UpdateAccountOptions) error {
	_, err := s.sq.
		Update(TABLE_NAME).
		Set(COLUMN_EMAIL, ops.Email).
		Where(squirrel.Eq{COLUMN_ID: id}).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("dbaccountstorage.Storage.Update: %w", err)
	}
	return nil
}
