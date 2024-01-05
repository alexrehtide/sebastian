package remoteaccountstorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) buildWhere(ops model.ReadRemoteAccountOptions) squirrel.Eq {
	where := squirrel.Eq{}
	if ops.ID != 0 {
		where[COLUMN_ID] = ops.ID
	}
	if ops.AccountID != 0 {
		where[COLUMN_ACCOUNT_ID] = ops.AccountID
	}
	if ops.RemoteID != "" {
		where[COLUMN_REMOTE_ID] = ops.RemoteID
	}
	if ops.RemoteEmail != "" {
		where[COLUMN_REMOTE_EMAIL] = ops.RemoteEmail
	}
	if ops.Platform != "" {
		where[COLUMN_PLATFORM] = ops.Platform
	}
	return where
}
