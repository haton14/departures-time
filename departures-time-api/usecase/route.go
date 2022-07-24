package usecase

import (
	"fmt"

	"github.com/haton14/departures-time/departures-time-api/external"
)

type Route interface {
	GetRouting(from, to external.StationCode) (string, error)
}

type route struct {
	exspert external.Exspert
}

func NewRoute(exspert external.Exspert) Route {
	return route{
		exspert: exspert,
	}
}

func (r route) GetRouting(from, to external.StationCode) (string, error) {
	routingURL, err := r.exspert.GetRoutingURL(from, to)
	if err != nil {
		return "", fmt.Errorf("Exspert.GetRoutingURL(): %w", err)
	}
	return routingURL, nil
}
