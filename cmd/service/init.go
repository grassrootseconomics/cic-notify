package main

import (
	"context"
	"strings"

	"github.com/grassrootseconomics/celoutils"
	"github.com/grassrootseconomics/cic-custodial/pkg/logg"
	"github.com/grassrootseconomics/cic-custodial/pkg/redis"
	"github.com/grassrootseconomics/cic-notify/internal/store"
	"github.com/grassrootseconomics/cic-notify/internal/sub"
	"github.com/grassrootseconomics/cic-notify/internal/tasker"
	"github.com/knadh/goyesql/v2"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/nats-io/nats.go"
	"github.com/zerodha/logf"
)

// Load logger.
func initLogger() logf.Logger {
	loggOpts := logg.LoggOpts{}

	if debugFlag {
		loggOpts.Color = true
		loggOpts.Caller = true
		loggOpts.Debug = true
	}

	return logg.NewLogg(loggOpts)
}

// Load config file.
func initConfig() *koanf.Koanf {
	var (
		ko = koanf.New(".")
	)

	confFile := file.Provider(confFlag)
	if err := ko.Load(confFile, toml.Parser()); err != nil {
		lo.Fatal("init: could not load config file", "error", err)
	}

	if err := ko.Load(env.Provider("NOTIFY_", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, "NOTIFY_")), "__", ".")
	}), nil); err != nil {
		lo.Fatal("init: could not override config from env vars", "error", err)
	}

	if debugFlag {
		ko.Print()
	}

	return ko
}

// Load Celo chain provider.
func initCeloProvider() *celoutils.Provider {
	providerOpts := celoutils.ProviderOpts{
		RpcEndpoint: ko.MustString("chain.rpc_endpoint"),
	}

	if ko.Bool("chain.testnet") {
		providerOpts.ChainId = celoutils.TestnetChainId
	} else {
		providerOpts.ChainId = celoutils.MainnetChainId
	}

	provider, err := celoutils.NewProvider(providerOpts)
	if err != nil {
		lo.Fatal("init: critical error loading chain provider", "error", err)
	}

	return provider
}

// Load separate redis connection for the tasker on a reserved db namespace.
func initAsynqRedisPool() *redis.RedisPool {
	poolOpts := redis.RedisPoolOpts{
		DSN:          ko.MustString("asynq.dsn"),
		MinIdleConns: ko.MustInt("redis.min_idle_conn"),
	}

	pool, err := redis.NewRedisPool(context.Background(), poolOpts)
	if err != nil {
		lo.Fatal("init: critical error connecting to asynq redis db", "error", err)
	}

	return pool
}

// Load tasker client.
func initTaskerClient(redisPool *redis.RedisPool) *tasker.TaskerClient {
	return tasker.NewTaskerClient(tasker.TaskerClientOpts{
		RedisPool: redisPool,
	})
}

func initQueries(queriesPath string) goyesql.Queries {
	queries, err := goyesql.ParseFile(queriesPath)
	if err != nil {
		lo.Fatal("init: could not load queries file", "error", err)
	}

	return queries
}

// Load Postgres store.
func initPgStore() store.Store {
	store, err := store.NewPgStore(store.Opts{
		DSN:                  ko.MustString("postgres.dsn"),
		MigrationsFolderPath: migrationsFolderFlag,
		QueriesFolderPath:    queriesFlag,
	})
	if err != nil {
		lo.Fatal("init: critical error loading Postgres store", "error", err)
	}

	return store
}

// Init JetStream context for both pub/sub.
func initJetStream() (*nats.Conn, nats.JetStreamContext) {
	natsConn, err := nats.Connect(ko.MustString("jetstream.endpoint"))
	if err != nil {
		lo.Fatal("init: critical error connecting to NATS", "error", err)
	}

	js, err := natsConn.JetStream()
	if err != nil {
		lo.Fatal("init: bad JetStream opts", "error", err)

	}

	return natsConn, js
}

func initSub(natsConn *nats.Conn, jsCtx nats.JetStreamContext, taskerClient *tasker.TaskerClient) *sub.Sub {
	sub, err := sub.NewSub(sub.SubOpts{
		JsCtx:        jsCtx,
		Logg:         lo,
		NatsConn:     natsConn,
		TaskerClient: taskerClient,
	})
	if err != nil {
		lo.Fatal("init: critical error bootstrapping sub", "error", err)
	}

	return sub
}
