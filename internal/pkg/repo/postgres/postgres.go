package postgres

import (
	"archv1/internal/pkg/config"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	*bun.DB
}

func NewDB(cfg *config.Config) *DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PostgresUser,
		cfg.PostgresPWD,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
		cfg.PostgresSSLMode,
	)
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqlDB, pgdialect.New())

	return &DB{
		db,
	}
}

func (d *DB) Delete(ctx context.Context, userId int, id int, table string) error {
	_, err := d.NewUpdate().
		Table(table).
		Set("deleted_at = now()").
		Set("deleted_by = ?", userId).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetOffset(page, limit *int) int {
	if page != nil && limit != nil {
		return ((*page) - 1) * (*limit)
	}
	return 0
}

func (d *DB) GenPagination(offset, page, limit *int) string {
	if page != nil && offset != nil {
		query := fmt.Sprintf(` LIMIT %d OFFSET %d`, *limit, *offset)
		return query

	} else if page != nil && limit != nil {
		offset := *limit * (*page - 1)
		query := fmt.Sprintf(` LIMIT %d OFFSET %d`, *limit, offset)
		return query
	}

	return ""
}
