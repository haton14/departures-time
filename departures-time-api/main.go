package main

import (
	"log"
	"net/http"
	"os"

	custom_middleware "github.com/haton14/departures-time/departures-time-api/middleware"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	log.Println("Stating Server")
	e := echo.New()
	e.Validator = custom_middleware.NewValidator()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet},
	}))
	injector := NewInjector()

	// routers
	v1 := e.Group("v1")
	v1.GET("/nearby-stations", injector.nearbyStationsHandler.Get)
	v1.GET("/destinations", injector.destinationHandler.Get)
	v1.GET("/routes", injector.routeHandler.Get)
	e.GET("/", func(c echo.Context) error {
		type health struct {
			Status string `json:"status"`
		}
		h := health{
			Status: "ok",
		}
		return c.JSON(http.StatusOK, h)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	if err := e.Start(":" + port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
