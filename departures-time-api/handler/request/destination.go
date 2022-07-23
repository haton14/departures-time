package request

import (
	"fmt"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/labstack/echo/v5"
)

type DestinationGet struct {
	Name vo.StationName
}

type destinationGetBindObject struct {
	Name *string `query:"name" validate:"required"`
}

func NewDestinationGet(c echo.Context) (*DestinationGet, error) {
	bindObject := &destinationGetBindObject{}
	if err := c.Bind(bindObject); err != nil {
		return nil, fmt.Errorf("echo.Context.Bind(): %s", err)
	}
	if err := c.Validate(bindObject); err != nil {
		return nil, fmt.Errorf("echo.Context.Validate(): %s", err)
	}
	name, err := vo.NewStationName(*bindObject.Name)
	if err != nil {
		return nil, fmt.Errorf("NewStationName() Name %s: %s", *bindObject.Name, err)
	}

	return &DestinationGet{
		Name: *name,
	}, nil
}