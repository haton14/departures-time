package main

import (
	"os"

	"github.com/haton14/departures-time/departures-time-api/external"
	"github.com/haton14/departures-time/departures-time-api/handler"
	"github.com/haton14/departures-time/departures-time-api/repository"
	"github.com/haton14/departures-time/departures-time-api/usecase"
)

type injector struct {
	nearbyStationsHandler handler.NearbyStations
}

func NewInjector() injector {
	// external
	ExspertExternal := external.NewExspert(os.Getenv("EKISPERT_API_PATH"), os.Getenv("EKISPERT_API_KEY"))
	NearRestApiExternal := external.NewNeaRestApi(os.Getenv("NEA_REST_API_PATH"))

	// repository
	NearbyStationDetailRepository := repository.NewNearbyStationDetail(ExspertExternal)
	NearbyStationsRepository := repository.NewNearbyStations(NearRestApiExternal)

	// usecase
	NearbyStationsUsecase := usecase.NewNearbyStations(NearbyStationsRepository, NearbyStationDetailRepository)

	// handler
	NearbyStationsHandler := handler.NewNearbyStations(NearbyStationsUsecase)

	return injector{
		nearbyStationsHandler: NearbyStationsHandler,
	}
}
