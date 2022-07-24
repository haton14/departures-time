package main

import (
	"log"
	"net/http"
	"os"

	"github.com/haton14/departures-time/departures-time-api/middleware"
	"github.com/labstack/echo/v5"
)

func main() {
	log.Println("Stating Server")
	e := echo.New()
	e.Validator = middleware.NewValidator()
	injector := NewInjector()

	// routers
	v1 := e.Group("v1")
	v1.GET("/nearby-stations", injector.nearbyStationsHandler.Get)
	v1.GET("/destinations", injector.destinationHandler.Get)
	v1.GET("/routes", injector.routeHandler.Get)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	if err := e.Start(":" + port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
