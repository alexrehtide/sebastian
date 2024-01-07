package registrationstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/alexrehtide/sebastian/model"
)

func (s *Storage) Create(ctx context.Context, ops model.CreateRegistrationOptions) (id uint, err error) {
	err = s.sq.
		Insert(TABLE_NAME).
		Columns(
			COLUMN_EMAIL,
			COLUMN_USERNAME,
			COLUMN_PASSWORD,
			COLUMN_VERIFICATION_CODE,
			COLUMN_CREATED_AT,
		).
		Values(
			ops.Email,
			ops.Username,
			ops.Password,
			ops.VerificationCode,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", COLUMN_ID)).
		ScanContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("registrationstorage.Storage.Create: %w", err)
	}
	return id, nil
}
