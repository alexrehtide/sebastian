package logstorage

import (
	"context"
	"fmt"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Create(ctx context.Context, ops model.CreateLogOptions) (id uint, err error) {
	var accountID *uint
	var sessionID *uint

	if ops.AccountID != 0 {
		accountID = &ops.AccountID
	}
	if ops.SessionID != 0 {
		sessionID = &ops.SessionID
	}

	err = s.sq.
		Insert(TABLE_NAME).
		Columns(
			COLUMN_LEVEL,
			COLUMN_ACCOUNT_ID,
			COLUMN_SESSION_ID,
			COLUMN_MESSAGE,
			COLUMN_DATA,
			COLUMN_CREATED_AT,
		).
		Values(
			ops.Level,
			accountID,
			sessionID,
			ops.Message,
			ops.Data,
			ops.CreatedAt,
		).
		Suffix(fmt.Sprintf("RETURNING %s", COLUMN_ID)).
		ScanContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("logstorage.Storage.Create: %w", err)
	}
	return id, nil
}
