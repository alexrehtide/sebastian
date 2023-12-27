package dbaccountrolestorage

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

func (s *Storage) Count(ctx context.Context, ops model.ReadAccountRoleOptions) (count int, err error) {
	err = s.sq.
		Select("count(*)").
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		QueryRowContext(ctx).
		Scan(count)
	if err != nil {
		return 0, fmt.Errorf("dbaccountrolestorage.Storage.Count: %w", err)
	}
	return
}

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

func (s *Storage) Delete(ctx context.Context, id uint) error {
	_, err := s.sq.
		Delete(TABLE_NAME).
		Where(squirrel.Eq{COLUMN_ID: id}).
		QueryContext(ctx)
	if err != nil {
		return fmt.Errorf("dbaccountrolestorage.Storage.Delete: %w", err)
	}
	return nil
}

func (s *Storage) Read(ctx context.Context, ops model.ReadAccountRoleOptions, pgOps model.PaginationOptions) (rows []model.AccountRole, err error) {
	sql, args, err := s.sq.
		Select(COLUMN_ID, COLUMN_ACCOUNT_ID, COLUMN_ROLE).
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		Limit(uint64(pgOps.Limit)).
		Offset(uint64(pgOps.Offset)).
		ToSql()
	if err != nil {
		return []model.AccountRole{}, fmt.Errorf("dbaccountrolestorage.Storage.scanRows: %w", err)
	}
	if err := s.db.SelectContext(ctx, rows, sql, args...); err != nil {
		return []model.AccountRole{}, fmt.Errorf("dbaccountrolestorage.Storage.scanRows: %w", err)
	}
	return
}

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
