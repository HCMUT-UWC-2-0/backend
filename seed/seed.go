package seed

import (
	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
)

// }
type Seed struct {
	// db      *sql.DB
	store      db.Store
	// source  []SeedSource
}

func NewSeed(store *db.Store) (*Seed, error) {
	seed := &Seed{
		store:      *store,
		// source: source,
	}
	return seed, nil
}

func (seed *Seed) Run() error {
	err := seed.runBackOfficersSeed()
	if err != nil {
		return err
	}

	err = seed.runWorkersSeed()
	if err != nil {
		return err
	}

	err = seed.runVehicleSeed()
	if err != nil {
		return err
	}

	err = seed.runMCPSeed()
	if err != nil {
		return err
	}

	err = seed.runRouteSeed()
	if err != nil {
		return err
	}

	return nil
}