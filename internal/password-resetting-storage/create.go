package passwordresettingstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Create(ctx context.Context, ops model.CreatePasswordResettingOptions) (id uint, err error) {
	err = s.sq.
		Insert(TABLE_NAME).
		Columns(
			COLUMN_ACCOUNT_ID,
			COLUMN_RESETTING_CODE,
			COLUMN_CREATED_AT,
		).
		Values(
			ops.AccountID,
			ops.ResettingCode,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", COLUMN_ID)).
		ScanContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("passwordresettingstorage.Storage.Create: %w", err)
	}
	return id, nil
}
