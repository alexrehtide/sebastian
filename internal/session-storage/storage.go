package dbsessionstorage

import (
	"context"
	"fmt"
	"time"

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

func (s *Storage) Count(ctx context.Context, ops model.ReadSessionOptions) (count int, err error) {
	err = s.sq.
		Select("count(*)").
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		QueryRowContext(ctx).
		Scan(count)
	if err != nil {
		return 0, fmt.Errorf("dbsessionstorage.Storage.Count: %w", err)
	}
	return
}

func (s *Storage) Create(ctx context.Context, ops model.CreateSessionOptions) (id uint, err error) {
	err = s.sq.
		Insert(TABLE_NAME).
		Columns(COLUMN_ACCOUNT_ID, COLUMN_ACCESS_TOKEN, COLUMN_REFRESH_TOKEN, COLUMN_CREATED_AT, COLUMN_UPDATED_AT).
		Values(ops.AccountID, ops.AccessToken, ops.RefreshToken, time.Now(), time.Now()).
		Suffix(fmt.Sprintf("RETURNING %s", COLUMN_ID)).
		ScanContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("dbsessionstorage.Storage.Create: %w", err)
	}
	return id, nil
}

func (s *Storage) Delete(ctx context.Context, id uint) error {
	_, err := s.sq.
		Delete(TABLE_NAME).
		Where(squirrel.Eq{COLUMN_ID: id}).
		QueryContext(ctx)
	if err != nil {
		return fmt.Errorf("dbsessionstorage.Storage.Delete: %w", err)
	}
	return nil
}

func (s *Storage) Read(ctx context.Context, ops model.ReadSessionOptions, pgOps model.PaginationOptions) (rows []model.Session, err error) {
	sql, args, err := s.sq.
		Select(COLUMN_ID, COLUMN_ACCOUNT_ID, COLUMN_ACCESS_TOKEN, COLUMN_REFRESH_TOKEN, COLUMN_CREATED_AT, COLUMN_UPDATED_AT).
		From(TABLE_NAME).
		Where(s.buildWhere(ops)).
		Limit(uint64(pgOps.Limit)).
		Offset(uint64(pgOps.Offset)).
		ToSql()
	if err != nil {
		return []model.Session{}, fmt.Errorf("dbsessionstorage.Storage.Read: %w", err)
	}
	if err := s.db.SelectContext(ctx, rows, sql, args...); err != nil {
		return []model.Session{}, fmt.Errorf("dbsessionstorage.Storage.Read: %w", err)
	}
	return
}

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

func (s *Storage) Update(ctx context.Context, id uint, ops model.UpdateSessionOptions) error {
	_, err := s.sq.
		Update(TABLE_NAME).
		Set(COLUMN_ACCESS_TOKEN, ops.AccessToken).
		Set(COLUMN_REFRESH_TOKEN, ops.RefreshToken).
		Set(COLUMN_UPDATED_AT, time.Now()).
		Where(squirrel.Eq{COLUMN_ID: id}).
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("dbsessionstorage.Storage.Update: %w", err)
	}
	return nil
}
