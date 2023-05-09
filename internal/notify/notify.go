package notify

import (
	"github.com/Khan/genqlient/graphql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang-module/carbon/v2"
	"github.com/grassrootseconomics/celoutils"
	hasura "github.com/grassrootseconomics/cic-notify/internal/graphql"
	"github.com/grassrootseconomics/cic-notify/internal/store"
	"github.com/grassrootseconomics/cic-notify/internal/tasker"
	"github.com/grassrootseconomics/cic-notify/internal/template"
	"github.com/kamikazechaser/africastalking"
	"github.com/redis/go-redis/v9"
	"github.com/zerodha/logf"
)

type (
	Opts struct {
		AtApiKey          string
		AtUsername        string
		AtSandbox         bool
		AtShortCode       string
		CeloProvider      *celoutils.Provider
		HasuraAdminSecret string
		HasuraEndpoint    string
		Logg              logf.Logger
		RedisClient       *redis.Client
		Store             store.Store
		TaskerClient      *tasker.TaskerClient
		TgBotToken        string
		Timezone          string
	}

	Notify struct {
		AtClient          *africastalking.AtClient
		AtShortCode       string
		CeloProvider      *celoutils.Provider
		GraphQLClient     graphql.Client
		Logg              logf.Logger
		RedisClient       *redis.Client
		Store             store.Store
		TaskerClient      *tasker.TaskerClient
		TgClient          *tgbotapi.BotAPI
		Timezone          string
		TxNotifyTemplates *template.TxNotifyTemplates
	}
)

func New(o Opts) (*Notify, error) {
	notifyContainer := Notify{
		AtClient:          africastalking.New(o.AtApiKey, o.AtUsername, o.AtSandbox),
		AtShortCode:       o.AtShortCode,
		CeloProvider:      o.CeloProvider,
		GraphQLClient:     hasura.NewHasuraGraphQLClient(o.HasuraAdminSecret, o.HasuraEndpoint),
		Logg:              o.Logg,
		RedisClient:       o.RedisClient,
		Store:             o.Store,
		TaskerClient:      o.TaskerClient,
		Timezone:          carbon.Moscow,
		TxNotifyTemplates: template.LoadTemplates(),
	}

	bot, err := tgbotapi.NewBotAPI(o.TgBotToken)
	if err != nil {
		return nil, err
	}
	notifyContainer.TgClient = bot

	return &notifyContainer, nil
}
