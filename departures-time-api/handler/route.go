package handler

import (
	"log"
	"net/http"

	"github.com/haton14/departures-time/departures-time-api/handler/request"
	"github.com/haton14/departures-time/departures-time-api/handler/response"
	"github.com/haton14/departures-time/departures-time-api/usecase"
	"github.com/labstack/echo/v5"
)

type Route struct {
	routeUsecase usecase.Route
}

func NewRoute(u usecase.Route) Route {
	return Route{
		routeUsecase: u,
	}
}

func (n Route) Get(c echo.Context) error {
	req, err := request.NewRouteGet(c)
	if err != nil {
		log.Printf("[Error] NewRouteGet(): %s", err)
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "リクエストが不正"})
	}
	result, err := n.routeUsecase.GetRouting(req.From, req.To)
	if err != nil {
		log.Printf("[Error] Route.GetRouting(): %s", err)
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "サーバ内のエラー"})
	}
	res := response.NewRouteGet(result)
	return c.JSON(http.StatusOK, res)
}
