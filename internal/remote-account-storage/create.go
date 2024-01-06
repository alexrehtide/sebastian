package remoteaccountstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Create(ctx context.Context, ops model.CreateRemoteAccountOptions) (id uint, err error) {
	var accountID *uint
	if ops.AccountID != 0 {
		accountID = &ops.AccountID
	}
	err = s.sq.
		Insert(TABLE_NAME).
		Columns(COLUMN_ACCOUNT_ID, COLUMN_REMOTE_ID, COLUMN_REMOTE_EMAIL, COLUMN_PLATFORM).
		Values(accountID, ops.RemoteID, ops.RemoteEmail, ops.Platform).
		Suffix(fmt.Sprintf("RETURNING %s", COLUMN_ID)).
		ScanContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("remoteaccountstorage.Storage.Create: %w", err)
	}
	return id, nil
}
