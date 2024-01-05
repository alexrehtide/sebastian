package accountstorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) buildWhere(ops model.ReadAccountOptions) squirrel.Eq {
	where := squirrel.Eq{}
	if ops.ID != 0 {
		where[COLUMN_ID] = ops.ID
	}
	if ops.Email != "" {
		where[COLUMN_EMAIL] = ops.Email
	}
	if ops.Username != "" {
		where[COLUMN_USERNAME] = ops.Username
	}
	return where
}
