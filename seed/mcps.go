package seed

import (
	"context"
	"fmt"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
)

func (seed *Seed) runMCPSeed() error {

	mcps := []db.CreateMCPParams{
		{
            Location: "Binh Thanh",
            Capacity: "3 tons",
        },
        {
            Location: "District 1",
            Capacity: "2 tons",
        },
        {
            Location: "Thu Duc",
            Capacity: "4 tons",
        },
        {
            Location: "District 7",
            Capacity: "1 ton",
        },
        {
            Location: "Go Vap",
            Capacity: "2.5 tons",
        },
        {
            Location: "Tan Binh",
            Capacity: "2 tons",
        },
        {
            Location: "Phu Nhuan",
            Capacity: "1.5 tons",
        },
        {
            Location: "District 3",
            Capacity: "3 tons",
        },
        {
            Location: "Binh Tan",
            Capacity: "2.5 tons",
        },
        {
            Location: "District 10",
            Capacity: "1 ton",
        },
	}

	// Insert the backofficers into the database
	for _, b := range mcps {
		seed.store.CreateMCP(context.Background(), b)
	}

	fmt.Println("MCPs seed completed.")

	return nil
}
