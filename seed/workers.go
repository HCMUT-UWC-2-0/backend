package seed

import (
	"context"
	"fmt"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	"github.com/HCMUT-UWC-2-0/backend/util"
)

func (seed *Seed) runWorkersSeed() error {

	workers := make([]db.CreateWorkerParams, 10)
	for i := 0; i < 10; i++ {
		workers[i] = util.RandomWorker(i)
	}

	// Insert the backofficers into the database
	for _, b := range workers {
		seed.store.CreateWorker(context.Background(), b)
	}

	fmt.Println("Back officers seed completed.")

	return nil
}
