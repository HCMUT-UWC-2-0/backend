package seed

import (
	"database/sql"
)

// type SeedSource struct {
// 	filePath string
// 	dbName 	string

// }
type Seed struct {
	db      *sql.DB
	// source  []SeedSource
}

func NewSeed(db *sql.DB) (*Seed, error) {
	seed := &Seed{
		db:      db,
		// source: source,
	}
	return seed, nil
}

func (seed *Seed) Run() error {
	err := seed.runBackOfficersSeed()
	if err != nil {
		return err
	}

	return nil
}