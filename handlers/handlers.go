package handlers

import (
	"errors"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mweegram/cti_api/db"
	"github.com/mweegram/cti_api/logic"
)

func API_Health(c echo.Context) error {
	if _, err := db.Database_Connect(); err != nil {
		return c.JSON(http.StatusInternalServerError, "{'Database Connection': 'Down'}")
	}
	return c.JSON(http.StatusOK, "{'Database Connection': 'Up'}")
}

func Create_Indicator_Handler(c echo.Context) error {
	new_indicator := db.Indicator{
		Type:    c.FormValue("type"),
		Value:   c.FormValue("value"),
		Comment: c.FormValue("comment"),
		Date:    time.Now().Format("2006-01-02"),
	}

	actor, err := strconv.Atoi(c.FormValue("actor"))
	if err != nil {
		actor = 1
	}
	new_indicator.Actor = actor

	ACCEPTABLE_TYPES := []string{"filehash", "ipaddress", "tactic", "cve", "email", "username", "hostname"}

	if !slices.Contains(ACCEPTABLE_TYPES, new_indicator.Type) {
		return errors.New("invalid indicator type")
	}

	if new_indicator.Value == "" {
		return errors.New("invalid idicator value")
	}

	err = logic.Create_Indicator(new_indicator)
	if err != nil {
		return err
	}

	return c.String(http.StatusAccepted, "New Indicator Created.")
}
