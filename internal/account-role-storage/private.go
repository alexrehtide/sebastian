package dbaccountrolestorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) buildWhere(ops model.ReadAccountRoleOptions) squirrel.Eq {
	where := squirrel.Eq{}
	if ops.ID != 0 {
		where[COLUMN_ID] = ops.ID
	}
	if ops.AccountID != 0 {
		where[COLUMN_ACCOUNT_ID] = ops.AccountID
	}
	if ops.Role != "" {
		where[COLUMN_ROLE] = ops.Role
	}
	return where
}
