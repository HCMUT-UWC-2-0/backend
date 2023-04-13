package seed

import (
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

	query := `INSERT INTO "BackOfficers" (email, ssn, hashed_password, name, phone, age, gender, date_of_birth, place_of_birth) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	backOfficers := []db.CreateBackOfficersParams{
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
        _, err := seed.db.Exec(query, b.Email, b.Ssn, b.HashedPassword, b.Name, b.Phone, b.Age, b.Gender, b.DateOfBirth, b.PlaceOfBirth)
        if err != nil {
            return err
        }
    }

	fmt.Println("Back officers seed completed.")

	return nil
}
