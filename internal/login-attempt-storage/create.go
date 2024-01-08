package loginattemptstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Create(ctx context.Context, ops model.CreateLoginAttemptOptions) (id uint, err error) {
	err = s.sq.
		Insert(TABLE_NAME).
		Columns(COLUMN_IP, COLUMN_COUNT, COLUMN_LAST_FAILED).
		Values(ops.IP, ops.Count, ops.LastFailed).
		Suffix(fmt.Sprintf("RETURNING %s", COLUMN_ID)).
		RunWith(s.getter.DefaultTrOrDB(ctx, s.db)).
		ScanContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("loginattemptstorage.Storage.Create: %w", err)
	}
	return id, nil
}
