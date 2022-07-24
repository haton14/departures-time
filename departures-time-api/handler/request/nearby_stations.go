package request

import (
	"fmt"
	"math"

	"github.com/haton14/departures-time/departures-time-api/domain/vo"
	"github.com/labstack/echo/v5"
)

type NearbyStationsGet struct {
	Longitude vo.Longitude
	Latitude  vo.Latitude
	Distance  vo.Distance
}

type nearbyStationsGetBindObject struct {
	Longitude *float64 `query:"longitude" validate:"required"`
	Latitude  *float64 `query:"latitude" validate:"required"`
	Distance  *int     `query:"distance"`
}

func NewNearbyStationsGet(c echo.Context) (*NearbyStationsGet, error) {
	bindObject := &nearbyStationsGetBindObject{}
	if err := c.Bind(bindObject); err != nil {
		return nil, fmt.Errorf("echo.Context.Bind(): %w", err)
	}
	if err := c.Validate(bindObject); err != nil {
		return nil, fmt.Errorf("echo.Context.Validate(): %w", err)
	}
	lo, err := vo.NewLongitude(*bindObject.Longitude)
	if err != nil {
		return nil, fmt.Errorf("NewLongitude() Longitude %f: %s", *bindObject.Longitude, err)
	}
	la, err := vo.NewLatitude(*bindObject.Latitude)
	if err != nil {
		return nil, fmt.Errorf("NewLatitude() Latitude %f: %s", *bindObject.Latitude, err)
	}
	if bindObject.Distance == nil {
		intMax := math.MaxInt
		bindObject.Distance = &intMax
	}
	distance, err := vo.NewDistance(*bindObject.Distance)
	if err != nil {
		return nil, fmt.Errorf("NewDistance() Distance %d: %s", *bindObject.Distance, err)
	}

	return &NearbyStationsGet{
		Longitude: *lo,
		Latitude:  *la,
		Distance:  *distance,
	}, nil
}
