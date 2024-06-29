package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mweegram/cti_api/db"
	"github.com/mweegram/cti_api/handlers"
)

func main() {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "Working")

		_, err := db.Database_Connect()
		if err != nil {
			fmt.Printf("%v", err)
		}

		return nil
	})

	app.GET("/health", handlers.API_Health)

	app.Start("localhost:5000")
}
