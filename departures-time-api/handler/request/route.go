package request

import (
	"fmt"

	"github.com/haton14/departures-time/departures-time-api/external"
	"github.com/labstack/echo/v5"
)

type RouteGet struct {
	From external.StationCode `query:"from" validate:"required"`
	To   external.StationCode `query:"to" validate:"required"`
}

func NewRouteGet(c echo.Context) (*RouteGet, error) {
	bindObject := &RouteGet{}
	if err := c.Bind(bindObject); err != nil {
		return nil, fmt.Errorf("echo.Context.Bind(): %w", err)
	}
	if err := c.Validate(bindObject); err != nil {
		return nil, fmt.Errorf("echo.Context.Validate(): %w", err)
	}

	return bindObject, nil
}
