package passwordresettingstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Count(ctx context.Context, ops model.ReadPasswordResettingOptions) (count int, err error) {
	err = s.sq.
		Select("count(*)").
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		RunWith(s.getter.DefaultTrOrDB(ctx, s.db)).
		QueryRowContext(ctx).
		Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("passwordresettingstorage.Storage.Count: %w", err)
	}
	return
}
