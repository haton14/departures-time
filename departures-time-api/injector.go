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
	destinationHandler    handler.Destination
}

func NewInjector() injector {
	// external
	ExspertExternal := external.NewExspert(os.Getenv("EKISPERT_API_PATH"), os.Getenv("EKISPERT_API_KEY"))
	NearRestApiExternal := external.NewNeaRestApi(os.Getenv("NEA_REST_API_PATH"))

	// repository
	NearbyStationDetailRepository := repository.NewNearbyStationDetail(ExspertExternal)
	NearbyStationsRepository := repository.NewNearbyStations(NearRestApiExternal)
	DestinationRepository := repository.NewDestination(ExspertExternal)

	// usecase
	NearbyStationsUsecase := usecase.NewNearbyStations(NearbyStationsRepository, NearbyStationDetailRepository)
	DestinationUsecase := usecase.NewDistination(DestinationRepository)

	// handler
	NearbyStationsHandler := handler.NewNearbyStations(NearbyStationsUsecase)
	DestinationHandler := handler.NewDestination(DestinationUsecase)

	return injector{
		nearbyStationsHandler: NearbyStationsHandler,
		destinationHandler:    DestinationHandler,
	}
}
