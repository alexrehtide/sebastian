package registrationformstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Count(ctx context.Context, ops model.ReadRegistrationFormOptions) (count int, err error) {
	err = s.sq.
		Select("count(*)").
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		QueryRowContext(ctx).
		Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("registrationformstorage.Storage.Count: %w", err)
	}
	return
}
