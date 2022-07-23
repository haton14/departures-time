package handler

import (
	"log"
	"net/http"

	"github.com/haton14/departures-time/departures-time-api/handler/request"
	"github.com/haton14/departures-time/departures-time-api/handler/response"
	"github.com/haton14/departures-time/departures-time-api/usecase"
	"github.com/labstack/echo/v5"
)

type Destination struct {
	destinationUsecase usecase.Destination
}

func NewDestination(u usecase.Destination) Destination {
	return Destination{destinationUsecase: u}
}

func (d Destination) Get(c echo.Context) error {
	req, err := request.NewDestinationGet(c)
	if err != nil {
		log.Printf("[Error] NewDestinationGet(): %s", err)
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "リクエストが不正"})
	}
	result, err := d.destinationUsecase.GetByName(req.Name)
	if err != nil {
		log.Printf("[Error] NearbyStations.GetByCoordinateAndDistance(): %s", err)
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "サーバ内のエラー"})
	}
	res := response.NewDestinationGet(result)
	return c.JSON(http.StatusOK, res)
}
