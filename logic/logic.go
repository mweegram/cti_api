package logic

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mweegram/cti_api/db"
)

func Create_Indicator(New_Indicator db.Indicator) error {
	database, err := db.Database_Connect()
	if err != nil {
		return err
	}

	if New_Indicator.Actor != 1 {
		err = database.QueryRow(context.Background(), "SELECT 1 FROM actors WHERE id = $1", New_Indicator.Actor).Scan()
		if err == nil {
			return errors.New("actor already exists")
		}

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

	var buffer int
	err = database.QueryRow(context.Background(), "SELECT 1 FROM actors WHERE name = $1", New_Actor.Name).Scan(&buffer)
	if err == nil {
		return errors.New("this group has already been added")
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

func Get_Indicator(Indicator_ID int) (db.Indicator_TextualActor, error) {
	database, err := db.Database_Connect()
	if err != nil {
		return db.Indicator_TextualActor{}, err
	}

	result, err := database.Query(context.Background(), "SELECT indicators.id as id,type,value,comment,date,actors.name as actor FROM indicators INNER JOIN actors ON indicators.actor = actors.id WHERE indicators.id = $1", Indicator_ID)
	if err != nil {
		return db.Indicator_TextualActor{}, err
	}
	indicator, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[db.Indicator_TextualActor])
	if err != nil {
		return db.Indicator_TextualActor{}, err
	}

	return indicator, nil
}

func Get_Actor(Actor_ID int) (db.Actor_Summary, error) {
	database, err := db.Database_Connect()
	if err != nil {
		return db.Actor_Summary{}, err
	}

	row, err := database.Query(context.Background(), "SELECT * FROM actors WHERE id = $1", Actor_ID)
	if err != nil {
		return db.Actor_Summary{}, err
	}

	actor_info, err := pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[db.Actor])
	if err != nil {
		return db.Actor_Summary{}, err
	}

	threat_actor_profile := db.Actor_Summary{
		ID:      actor_info.ID,
		Name:    actor_info.Name,
		Aliases: actor_info.Aliases,
	}

	rows, err := database.Query(context.Background(), "SELECT * FROM indicators WHERE actor = $1", Actor_ID)
	if err != nil {
		return db.Actor_Summary{}, err
	}

	threat_actor_profile.Indicators, err = pgx.CollectRows(rows, pgx.RowToStructByName[db.Indicator])
	if err != nil {
		return db.Actor_Summary{}, err
	}

	return threat_actor_profile, nil
}
