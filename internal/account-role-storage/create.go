package accountrolestorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Create(ctx context.Context, ops model.CreateAccountRoleOptions) (id uint, err error) {
	err = s.sq.
		Insert(TABLE_NAME).
		Columns(COLUMN_ACCOUNT_ID, COLUMN_ROLE).
		Values(ops.AccountID, ops.Role).
		Suffix(fmt.Sprintf("RETURNING %s", COLUMN_ID)).
		ScanContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("dbaccountrolestorage.Storage.Create: %w", err)
	}
	return id, nil
}
