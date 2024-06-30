package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mweegram/cti_api/handlers"
)

func main() {
	app := echo.New()

	app.GET("/health", handlers.API_Health)
	app.GET("/indicator/:id", handlers.Get_Indicator_Handler)
	app.GET("/actor/:id", handlers.Get_Actor_Handler)

	app.POST("/new_indicator", handlers.Create_Indicator_Handler)
	app.POST("/new_actor", handlers.Create_Actor_Handler)
	app.POST("/new_alias/:actor", handlers.Create_Alias_Handler)

	app.Start("localhost:5000")
}
