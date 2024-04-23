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
	profile, err := q.GetProfile(ctx, kthID);
	if err == pgx.ErrNoRows {
		profile, err = q.CreateProfile(ctx, kthID);
	}
	if err != nil {
		return nil, err
	}

	return &profile, nil
};

// A validatedProfile is a wrapper of 
// [github.com/datasektionen/midas/internal.db.Profile] allowing for updates
// to the data.
type validatedProfile struct {
	profile *db.Profile
}

// ValidateProfile wraps a [github.com/datasektionen/midas/internal.db.Profile]
// returning a 
// [github.com/datasektionen/midas/internal.profile.validatedProfile].
func ValidateProfile(p *db.Profile) validatedProfile {
	return validatedProfile{profile: p};
}

// UpdateBank updates the bank name field of a 
// [github.com/datasektionen/midas/internal.profile.validatedProfile] as well
// as the underlying database row. 
func (v *validatedProfile) UpdateBank(
	ctx context.Context,
	q *db.Queries, 
	bankName string,
) error {
	params := db.UpdateProfileBankParams{ID: v.profile.ID, Bank: bankName};

	if err := q.UpdateProfileBank(ctx, params); err != nil {
		return err
	}

	v.profile.Bank = bankName;

	return nil
}

// UpdateBankAccount updates the bank account number field of a 
// [github.com/datasektionen/midas/internal.profile.validatedProfile] as well
// as the underlying database row. 
func (v *validatedProfile) UpdateBankAccount(
	ctx context.Context,
	q *db.Queries, 
	bankAccountNumber string,
) error {
	params := db.UpdateProfileBankAccountNumberParams{
		ID: v.profile.ID,
		BankAccountNumber: bankAccountNumber,
	};

	if err := q.UpdateProfileBankAccountNumber(ctx, params); err != nil {
		return err
	}

	v.profile.BankAccountNumber = bankAccountNumber;

	return nil
}

// UpdateClearingNumber updates the clearing number field of a 
// [github.com/datasektionen/midas/internal.profile.validatedProfile] as well
// as the underlying database row. 
func (v *validatedProfile) UpdateClearingNumber(
	ctx context.Context,
	q *db.Queries, 
	clearingNumber string,
) error {
	params := db.UpdateProfileClearingNumberParams{
		ID: v.profile.ID, 
		ClearingNumber: clearingNumber,
	};

	if err := q.UpdateProfileClearingNumber(ctx, params); err != nil {
		return err
	}

	v.profile.ClearingNumber = clearingNumber;

	return nil
}
