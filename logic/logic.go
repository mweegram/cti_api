package logic

import (
	"context"

	"github.com/mweegram/cti_api/db"
)

func Create_Indicator(New_Indicator db.Indicator) error {
	database, err := db.Database_Connect()
	if err != nil {
		return err
	}

	const QUERY_STRING = "INSERT INTO indicators (type,value,comment,date,actor) VALUES ($1,$2,$3,$4,$5)"
	_, err = database.Exec(context.Background(), QUERY_STRING, New_Indicator.Type, New_Indicator.Value, New_Indicator.Comment, New_Indicator.Date, New_Indicator.Actor)
	if err != nil {
		return err
	}

	return nil
}

func Create_Actor(New_Actor db.Actor) error {
	database, err := db.Database_Connect()
	if err != nil {
		return err
	}

	const QUERY_STRING = "INSERT INTO actors (name,aliases) VALUES ($1,$2)"
	_, err = database.Exec(context.Background(), QUERY_STRING, New_Actor.Name, New_Actor.Aliases)
	if err != nil {
		return err
	}

	return nil
}

func Create_Alias(Actor_ID int, Alias string) error {
	database, err := db.Database_Connect()
	if err != nil {
		return err
	}

	const QUERY_STRING = "UPDATE actors SET aliases = array_append(aliases,$1) WHERE id = $2"
	_, err = database.Exec(context.Background(), QUERY_STRING, Alias, Actor_ID)
	if err != nil {
		return err
	}

	return nil
}
