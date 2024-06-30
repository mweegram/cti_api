package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mweegram/cti_api/handlers"
)

func main() {
	app := echo.New()

	app.GET("/health", handlers.API_Health)
	app.POST("/new_indicator", handlers.Create_Indicator_Handler)

	app.Start("localhost:5000")
}
