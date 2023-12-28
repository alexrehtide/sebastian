package dbaccountstorage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (s *Storage) Delete(ctx context.Context, id uint) error {
	_, err := s.sq.
		Delete(TABLE_NAME).
		Where(squirrel.Eq{COLUMN_ID: id}).
		QueryContext(ctx)
	if err != nil {
		return fmt.Errorf("dbaccountstorage.Storage.Delete: %w", err)
	}
	return nil
}
