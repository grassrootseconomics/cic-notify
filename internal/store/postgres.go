package store

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/tern/v2/migrate"
	"github.com/knadh/goyesql/v2"
	"github.com/zerodha/logf"
)

const (
	schemaTable = "schema_version"
)

type (
	queries struct {
		CreateAtReceipt string `query:"create-at-receipt"`
		CreateTgReceipt string `query:"create-tg-receipt"`
		SetAtDelivered  string `query:"set-at-delivered"`
	}

	PostgresStoreOpts struct {
		DSN                  string
		MigrationsFolderPath string
		Logg                 logf.Logger
		Queries              goyesql.Queries
	}

	PostgresStore struct {
		logg    logf.Logger
		pool    *pgxpool.Pool
		queries queries
	}
)

func NewPostgresStore(o PostgresStoreOpts) (Store, error) {
	postgresStore := &PostgresStore{
		logg: o.Logg,
	}

	if err := goyesql.ScanToStruct(&postgresStore.queries, o.Queries, nil); err != nil {
		return nil, fmt.Errorf("failed to scan queries %v", err)
	}

	parsedConfig, err := pgxpool.ParseConfig(o.DSN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbPool, err := pgxpool.NewWithConfig(ctx, parsedConfig)
	if err != nil {
		return nil, err
	}

	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	migrator, err := migrate.NewMigrator(ctx, conn.Conn(), schemaTable)
	if err != nil {
		return nil, err
	}

	if err := migrator.LoadMigrations(os.DirFS(o.MigrationsFolderPath)); err != nil {
		return nil, err
	}

	if err := migrator.Migrate(ctx); err != nil {
		return nil, err
	}

	postgresStore.pool = dbPool

	return postgresStore, nil
}

func (s *PostgresStore) CreateAtReceipt(ctx context.Context, statusCode uint, messageId string) error {
	_, err := s.pool.Exec(ctx, s.queries.CreateAtReceipt, statusCode, messageId)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) CreateTgReceipt(ctx context.Context, messageId int) error {
	_, err := s.pool.Exec(ctx, s.queries.CreateTgReceipt, messageId)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) SetAtDelivered(ctx context.Context, messageId string) error {
	_, err := s.pool.Exec(ctx, s.queries.SetAtDelivered, messageId)
	if err != nil {
		return err
	}

	return nil
}
