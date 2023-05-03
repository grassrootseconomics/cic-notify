package graphql

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
)

type authedTransport struct {
	adminSecret string
	wrapped     http.RoundTripper
}

func NewHasuraGraphQLClient(AdminSecret string, Endpoint string) graphql.Client {
	httpClient := http.Client{
		Timeout: time.Second * 5,
		Transport: &authedTransport{
			adminSecret: AdminSecret,
			wrapped:     http.DefaultTransport,
		},
	}
	return graphql.NewClient(fmt.Sprintf("%s/v1/graphql", Endpoint), &httpClient)
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("x-hasura-admin-secret", t.adminSecret)
	return t.wrapped.RoundTrip(req)
}
