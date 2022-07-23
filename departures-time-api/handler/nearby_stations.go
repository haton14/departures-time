package handler

import (
	"log"
	"net/http"

	"github.com/haton14/departures-time/departures-time-api/handler/request"
	"github.com/haton14/departures-time/departures-time-api/handler/response"
	"github.com/haton14/departures-time/departures-time-api/usecase"
	"github.com/labstack/echo/v5"
)

type NearbyStations struct {
	nearbyStationsUsecase usecase.NearbyStations
}

func NewNearbyStations(u usecase.NearbyStations) NearbyStations {
	return NearbyStations{nearbyStationsUsecase: u}
}

func (n NearbyStations) Get(c echo.Context) error {
	req, err := request.NewNearbyStationsGet(c)
	if err != nil {
		log.Printf("[Error] NewNearbyStationsGet(): %s", err)
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "リクエストが不正"})
	}
	result, err := n.nearbyStationsUsecase.GetByCoordinateAndDistance(req.Longitude, req.Latitude, req.Distance)
	if err != nil {
		log.Printf("[Error] NearbyStations.GetByCoordinateAndDistance(): %s", err)
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "サーバ内のエラー"})
	}
	res := response.NewNearbyStationsGet(result)
	return c.JSON(http.StatusOK, res)
}
