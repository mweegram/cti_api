package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mweegram/cti_api/db"
)

func API_Health(c echo.Context) error {
	if _, err := db.Database_Connect(); err != nil {
		return c.JSON(http.StatusInternalServerError, "{'Database Connection': 'Down'}")
	}
	return c.JSON(http.StatusOK, "{'Database Connection': 'Up'}")
}
