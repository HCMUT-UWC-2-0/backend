package seed

import (
	"context"
	"fmt"
	"time"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	"github.com/HCMUT-UWC-2-0/backend/util"
)

func (seed *Seed) runBackOfficersSeed() error {

	hashedPassword, err := util.HashPassword("secret")
	if err != nil {
		return err
	}


	backOfficers := []db.CreateBackOfficerParams{
		{
			Email:          "john.doe@example.com",
			Ssn:            "123-45-6789",
			HashedPassword: hashedPassword,
			Name:           "John Doe",
			Phone:          "555-1234",
			Age:            30,
			Gender:         db.GenderTypeMALE,
			DateOfBirth:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth:   "New York, NY",
		},
		{
			Email:          "jane.doe@example.com",
			Ssn:            "987-65-4321",
			HashedPassword: hashedPassword,
			Name:           "Jane Doe",
			Phone:          "555-5678",
			Age:            25,
			Gender:         db.GenderTypeFEMALE,
			DateOfBirth:    time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth:   "Los Angeles, CA",
		},
	}


	// Insert the backofficers into the database
    for _, b := range backOfficers {
        seed.store.CreateBackOfficer(context.Background(),b)
    }

	fmt.Println("Back officers seed completed.")

	return nil
}
