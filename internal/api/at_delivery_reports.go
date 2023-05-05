package api

import (
	"net/http"

	"github.com/grassrootseconomics/cic-notify/internal/store"
	"github.com/labstack/echo/v4"
)

// WIP
// Updates At receipt delivery status.
// https://developers.africastalking.com/docs/sms/notifications
func HandleAtDeliveryReport(s store.Store) func(echo.Context) error {
	return func(c echo.Context) error {
		var req struct {
			MessageId string `form:"id"`
			Status    string `form:"status"`
		}

		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid form data.")
		}

		if err := s.SetAtDelivered(c.Request().Context(), req.MessageId); err != nil {
			return err
		}

		return c.String(http.StatusOK, "Ok!")
	}
}
