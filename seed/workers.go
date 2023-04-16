package seed

import (
	"context"
	"fmt"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	"github.com/HCMUT-UWC-2-0/backend/util"
)

func (seed *Seed) runWorkersSeed() error {

	janitors := make([]db.CreateWorkerParams, 10)
	for i := 0; i < 10; i++ {
		janitors[i] = util.RandomWorker(i, db.WorkerTypeJANITOR)
	}

	// Insert the backofficers into the database
	for _, b := range janitors {
		worker, _ := seed.store.CreateWorker(context.Background(), b)
		seed.store.CreateWorkerStatus(context.Background(), int32(worker.ID))
		
	}


	collectors := make([]db.CreateWorkerParams, 10)
	for i := 0; i < 10; i++ {
		collectors[i] = util.RandomWorker(i+10, db.WorkerTypeCOLLECTOR)
		
	}
	
	// Insert the backofficers into the database
	for _, b := range collectors {
		worker, _ := seed.store.CreateWorker(context.Background(), b)
		seed.store.CreateWorkerStatus(context.Background(), int32(worker.ID))
	}
	fmt.Println("Janitors and collectors seeds completed.")

	return nil
}
