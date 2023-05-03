package notify

import (
	"github.com/Khan/genqlient/graphql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang-module/carbon/v2"
	"github.com/grassrootseconomics/celoutils"
	hasura "github.com/grassrootseconomics/cic-notify/internal/graphql"
	"github.com/grassrootseconomics/cic-notify/internal/store"
	"github.com/grassrootseconomics/cic-notify/internal/tasker"
	"github.com/grassrootseconomics/cic-notify/internal/templates"
	"github.com/kamikazechaser/africastalking"
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
		Store             store.Store
		TaskerClient      *tasker.TaskerClient
		TgClient          *tgbotapi.BotAPI
		Timezone          string
		TxNotifyTemplates *templates.TxNotifyTemplates
	}
)

func New(o Opts) (*Notify, error) {
	notifyContainer := Notify{
		AtClient:          africastalking.New(o.AtApiKey, o.AtUsername, o.AtSandbox),
		AtShortCode:       o.AtShortCode,
		CeloProvider:      o.CeloProvider,
		GraphQLClient:     hasura.NewHasuraGraphQLClient(o.HasuraAdminSecret, o.HasuraEndpoint),
		Store:             o.Store,
		TaskerClient:      o.TaskerClient,
		Timezone:          carbon.Moscow,
		TxNotifyTemplates: templates.LoadTemplates(),
	}

	bot, err := tgbotapi.NewBotAPI(o.TgBotToken)
	if err != nil {
		return nil, err
	}
	notifyContainer.TgClient = bot

	return &notifyContainer, nil
}
