package registrationformstorage

import (
	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) buildWhere(ops model.ReadRegistrationFormOptions) squirrel.Eq {
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
	if ops.VerificationCode != "" {
		where[COLUMN_VERIFICATION_CODE] = ops.VerificationCode
	}
	return where
}
