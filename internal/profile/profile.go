// Package profile implements functionality for accessing or modifying profiles
// in the system.
package profile

import (
	"context"

	"github.com/datasektionen/midas/internal/db"
	"github.com/jackc/pgx/v5"
)

// GetOrCreate returns a [github.com/datasektionen/midas/internal.db.Profile]
// either by fetching it from the database or creating it if it doesn't already
// exist. Profile is nil in case of error.
func GetOrCreate(
	ctx context.Context,
	q *db.Queries,
	kthID string,
) (*db.Profile, error) {
	profile, err := Get(ctx, q, kthID)
	if err == pgx.ErrNoRows {
		profile, err = Create(ctx, q, kthID);
	}
	if err != nil {
		return nil, err
	}

	return profile, nil
};

// Create returns a [github.com/datasektionen/midas/internal.db.Profile]
// by fetching it from the database. Profile is nil in case of error.
func Create(
	ctx context.Context,
	q *db.Queries,
	kthID string,
) (*db.Profile, error) {
	profile, err := q.CreateProfile(ctx, kthID);
	if err != nil {
		return nil, err
	}

	return &profile, nil
};

// Create returns a [github.com/datasektionen/midas/internal.db.Profile]
// by creating a new user with that kthID in the database. 
// Profile is nil in case of error.
func Get(
	ctx context.Context,
	q *db.Queries,
	kthID string,
) (*db.Profile, error) {
	profile, err := q.GetProfile(ctx, kthID);
	if err != nil {
		return nil, err
	}

	return &profile, nil
};
