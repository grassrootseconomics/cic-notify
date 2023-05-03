package main

import (
	"fmt"
	"net/http"

	"github.com/VictoriaMetrics/metrics"
	"github.com/grassrootseconomics/cic-custodial/pkg/util"
	"github.com/grassrootseconomics/cic-notify/internal/api"
	"github.com/grassrootseconomics/cic-notify/internal/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Bootstrap API server.
func initApiServer(store store.Store) *echo.Echo {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true
	server.HTTPErrorHandler = customHTTPErrorHandler

	server.Use(middleware.Recover())
	server.Use(middleware.BodyLimit("1M"))
	server.Use(middleware.ContextTimeout(util.SLATimeout))

	if ko.Bool("service.metrics") {
		server.GET("/metrics", func(c echo.Context) error {
			metrics.WritePrometheus(c.Response(), true)
			return nil
		})
	}

	webhookRoute := server.Group("/webhook")
	webhookRoute.POST(fmt.Sprint("/at/%s", ko.MustString("at.webhook_secret")), api.HandleAtDeliveryReport(store))

	return server
}

func customHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	if he, ok := err.(*echo.HTTPError); ok {
		var errorMsg string

		if m, ok := he.Message.(error); ok {
			errorMsg = m.Error()
		} else if m, ok := he.Message.(string); ok {
			errorMsg = m
		}

		c.String(he.Code, errorMsg)
		return
	}

	lo.Error("api: echo error", "path", c.Path(), "err", err)
	c.String(http.StatusInternalServerError, "Internal server error.")
}
