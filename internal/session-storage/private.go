package sessionstorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) buildWhere(ops model.ReadSessionOptions) squirrel.Eq {
	where := squirrel.Eq{}
	if ops.ID != 0 {
		where[COLUMN_ID] = ops.ID
	}
	if ops.AccountID != 0 {
		where[COLUMN_ACCOUNT_ID] = ops.AccountID
	}
	if ops.AccessToken != "" {
		where[COLUMN_ACCESS_TOKEN] = ops.AccessToken
	}
	if ops.RefreshToken != "" {
		where[COLUMN_REFRESH_TOKEN] = ops.RefreshToken
	}
	return where
}
