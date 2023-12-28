package accountstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Create(ctx context.Context, ops model.CreateAccountOptions) (id uint, err error) {
	err = s.sq.
		Insert(TABLE_NAME).
		Columns(COLUMN_EMAIL, COLUMN_PASSWORD_HASH).
		Values(ops.Email, ops.Password).
		Suffix(fmt.Sprintf("RETURNING %s", COLUMN_ID)).
		ScanContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("accountstorage.Storage.Create: %w", err)
	}
	return
}
