package main

import (
	"context"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"

	"github.com/datasektionen/midas/internal/db"
	"github.com/datasektionen/midas/internal/profile"
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
	created, err := profile.GetOrCreate(ctx, queries, "emilhul")
	if err != nil {
		return err
	}
	log.Println(created)

	// create authenticedProile
	ap := profile.ValidateProfile(created);

	// Update bank
	ap.UpdateBank(ctx, queries, "Handelsbanken");
	ap.UpdateBankAccount(ctx, queries, "1337420");
	ap.UpdateClearingNumber(ctx, queries, "666");
	log.Println(created)

	// get a profile
	fetched, err := profile.GetOrCreate(ctx, queries, "emilhul")
	if err != nil {
		return err
	}
	log.Println(fetched)

	// prints true (ap is a wrapper to a reference to created)
	log.Println(reflect.DeepEqual(fetched, created))

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
