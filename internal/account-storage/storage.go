package dbaccountstorage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexrehtide/sebastian/model"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
		sq: squirrel.StatementBuilder.RunWith(db).PlaceholderFormat(squirrel.Dollar),
	}
}

type Storage struct {
	db *sqlx.DB
	sq squirrel.StatementBuilderType
}

func (s *Storage) Count(ctx context.Context, ops model.ReadAccountOptions) (count int, err error) {
	err = s.sq.
		Select("count(*)").
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		QueryRowContext(ctx).
		Scan(count)
	if err != nil {
		return 0, fmt.Errorf("dbaccountstorage.Storage.Count: %w", err)
	}
	return
}

func (s *Storage) Create(ctx context.Context, ops model.CreateAccountOptions) (id uint, err error) {
	err = s.sq.
		Insert(TABLE_NAME).
		Columns(COLUMN_EMAIL, COLUMN_PASSWORD_HASH).
		Values(ops.Email, ops.Password).
		Suffix(fmt.Sprintf("RETURNING %s", COLUMN_ID)).
		ScanContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("dbaccountstorage.Storage.Create: %w", err)
	}
	return
}

func (s *Storage) Delete(ctx context.Context, id uint) error {
	_, err := s.sq.
		Delete(TABLE_NAME).
		Where(squirrel.Eq{COLUMN_ID: id}).
		QueryContext(ctx)
	if err != nil {
		return fmt.Errorf("dbaccountstorage.Storage.Delete: %w", err)
	}
	return nil
}

func (s *Storage) Read(ctx context.Context, ops model.ReadAccountOptions, pgOps model.PaginationOptions) (rows []model.Account, err error) {
	sql, args, err := s.sq.
		Select(COLUMN_ID, COLUMN_EMAIL, COLUMN_PASSWORD_HASH).
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		Limit(uint64(pgOps.Limit)).
		Offset(uint64(pgOps.Offset)).
		ToSql()
	if err != nil {
		return []model.Account{}, fmt.Errorf("dbaccountstorage.Storage.Read: %w", err)
	}
	if err := s.db.SelectContext(ctx, rows, sql, args...); err != nil {
		return []model.Account{}, fmt.Errorf("dbaccountstorage.Storage.Read: %w", err)
	}
	return
}

func (s *Storage) buildWhere(ops model.ReadAccountOptions) squirrel.Eq {
	where := squirrel.Eq{}
	if ops.ID != 0 {
		where[COLUMN_ID] = ops.ID
	}
	if ops.Email != "" {
		where[COLUMN_EMAIL] = ops.Email
	}
	return where
}

func (s *Storage) Update(ctx context.Context, id uint, ops model.UpdateAccountOptions) error {
	_, err := s.sq.
		Update(TABLE_NAME).
		Set(COLUMN_EMAIL, ops.Email).
		Where(squirrel.Eq{COLUMN_ID: id}).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("dbaccountstorage.Storage.Update: %w", err)
	}
	return nil
}
