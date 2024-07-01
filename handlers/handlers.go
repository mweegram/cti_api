package handlers

import (
	"context"
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

func Create_Actor_Handler(c echo.Context) error {
	new_actor := db.Actor{
		Name: c.FormValue("name"),
	}

	if new_actor.Name == "" {
		return errors.New("cannot have blank name")
	}

	err := logic.Create_Actor(new_actor)
	if err != nil {
		return err
	}

	return c.String(http.StatusAccepted, "New actor added.")
}

func Create_Alias_Handler(c echo.Context) error {
	threat_actor_str := c.Param("actor")
	if threat_actor_str == "" {
		return errors.New("invalid threat actor")
	}

	threat_actor, err := strconv.Atoi(threat_actor_str)
	if err != nil {
		return err
	}

	new_alias := c.FormValue("alias")
	if new_alias == "" {
		return errors.New("invalid alias")
	}

	database, err := db.Database_Connect()
	if err != nil {
		return err
	}

	var cur_aliases []string

	err = database.QueryRow(context.Background(), "SELECT aliases FROM actors WHERE id = $1", threat_actor).Scan(&cur_aliases)
	if err != nil {
		return err
	}

	if slices.Contains(cur_aliases, new_alias) {
		return errors.New("alias already exists")
	}

	err = logic.Create_Alias(threat_actor, new_alias)
	if err != nil {
		return err
	}

	return c.String(http.StatusAccepted, "Alias added.")
}

func Get_Indicator_Handler(c echo.Context) error {
	indicator_id_str := c.Param("id")
	if indicator_id_str == "" {
		return errors.New("invalid threat id")
	}

	indicator_id, err := strconv.Atoi(indicator_id_str)
	if err != nil {
		return err
	}

	indicator, err := logic.Get_Indicator(indicator_id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, indicator)
}

func Get_Actor_Handler(c echo.Context) error {
	actor_id_str := c.Param("id")
	if actor_id_str == "" {
		return errors.New("invalid threat id")
	}

	actor_id, err := strconv.Atoi(actor_id_str)
	if err != nil {
		return err
	}

	actor, err := logic.Get_Actor(actor_id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, actor)
}

func Get_AllActors_Handler(c echo.Context) error {
	threat_actors, err := logic.Get_All_Actors()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, threat_actors)
}
