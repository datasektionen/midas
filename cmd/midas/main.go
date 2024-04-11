package main

import (
    "context"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"

    "github.com/datasektionen/midas/internal/db"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=localhost password=midas user=midas dbname=midas sslmode=prefer")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	// list all profiles
	profiles, err := queries.ListProfiles(ctx)
	if err != nil {
		return err
	}
	log.Println(profiles)

	// create a profile
	insertedProfile, err := queries.CreateProfile(ctx, "emilhul")
	if err != nil {
		return err
	}
	log.Println(insertedProfile)

	// get the profile we just inserted
	fetchedProfile, err := queries.GetProfile(ctx, insertedProfile.KthID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedProfile, fetchedProfile))

	// list all profiles again
	profiles, err = queries.ListProfiles(ctx)
	if err != nil {
		return err
	}
	log.Println(profiles)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
