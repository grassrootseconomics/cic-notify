package store

import (
	"context"
	"fmt"
	"os"

	"github.com/grassrootseconomics/cic-custodial/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/tern/v2/migrate"
	"github.com/knadh/goyesql/v2"
)

const (
	schemaTable = "schema_version"
)

type (
	Store interface {
		CreateAtReceipt(context.Context, uint, string) error
		CreateTgReceipt(context.Context, int) error
		SetAtDelivered(context.Context, string) error
	}

	Opts struct {
		DSN                  string
		MigrationsFolderPath string
		QueriesFolderPath    string
	}

	PgStore struct {
		db      *pgxpool.Pool
		queries *queries
	}
	queries struct {
		CreateAtReceipt string `query:"create-at-receipt"`
		CreateTgReceipt string `query:"create-tg-receipt"`
		SetAtDelivered  string `query:"set-at-delivered"`
	}
)

func NewPgStore(o Opts) (Store, error) {
	parsedConfig, err := pgxpool.ParseConfig(o.DSN)
	if err != nil {
		return nil, err
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), parsedConfig)
	if err != nil {
		return nil, err
	}

	queries, err := loadQueries(o.QueriesFolderPath)
	if err != nil {
		return nil, err
	}

	if err := runMigrations(context.Background(), dbPool, o.MigrationsFolderPath); err != nil {
		return nil, err
	}

	return &PgStore{
		db:      dbPool,
		queries: queries,
	}, nil
}

func loadQueries(queriesPath string) (*queries, error) {
	parsedQueries, err := goyesql.ParseFile(queriesPath)
	if err != nil {
		return nil, err
	}

	loadedQueries := &queries{}

	if err := goyesql.ScanToStruct(loadedQueries, parsedQueries, nil); err != nil {
		return nil, fmt.Errorf("failed to scan queries %v", err)
	}

	return loadedQueries, nil
}

func runMigrations(ctx context.Context, dbPool *pgxpool.Pool, migrationsPath string) error {
	ctx, cancel := context.WithTimeout(ctx, util.SLATimeout)
	defer cancel()

	conn, err := dbPool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	migrator, err := migrate.NewMigrator(ctx, conn.Conn(), "schema_version")
	if err != nil {
		return err
	}

	if err := migrator.LoadMigrations(os.DirFS(migrationsPath)); err != nil {
		return err
	}

	if err := migrator.Migrate(ctx); err != nil {
		return err
	}

	return nil
}
