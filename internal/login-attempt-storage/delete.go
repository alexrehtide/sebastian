package loginattemptstorage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (s *Storage) Delete(ctx context.Context, ip string) error {
	_, err := s.sq.
		Delete(TABLE_NAME).
		Where(squirrel.Eq{COLUMN_IP: ip}).
		RunWith(s.getter.DefaultTrOrDB(ctx, s.db)).
		QueryContext(ctx)
	if err != nil {
		return fmt.Errorf("loginattemptstorage.Storage.Delete: %w", err)
	}
	return nil
}
