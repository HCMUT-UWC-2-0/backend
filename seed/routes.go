package seed

import (
	"context"
	"fmt"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
)

func (seed *Seed) runRouteSeed() error {

	routes := []db.CreateRouteParams{
		{
			StartLocation: "Ben Thanh",
			EndLocation:   "Cho Lon",
			Distance:      "6.5 km",
			EstimatedTime: "30 minutes",
		},
		{
			StartLocation: "Ben Thanh",
			EndLocation:   "An Suong",
			Distance:      "15.5 km",
			EstimatedTime: "1 hour",
		},
		{
			StartLocation: "Ben Thanh",
			EndLocation:   "Tan Binh",
			Distance:      "7.3 km",
			EstimatedTime: "40 minutes",
		},
		{
			StartLocation: "Ben Thanh",
			EndLocation:   "Nguyen Van Cu",
			Distance:      "5.2 km",
			EstimatedTime: "25 minutes",
		},
		{
			StartLocation: "Cho Lon",
			EndLocation:   "Phu Lam",
			Distance:      "12.5 km",
			EstimatedTime: "1 hour 15 minutes",
		},
		{
			StartLocation: "Cho Lon",
			EndLocation:   "Hiep Thanh",
			Distance:      "15.5 km",
			EstimatedTime: "1 hour 30 minutes",
		},
		{
			StartLocation: "An Suong",
			EndLocation:   "Hoc Mon",
			Distance:      "25 km",
			EstimatedTime: "2 hours",
		},
		{
			StartLocation: "Tan Binh",
			EndLocation:   "Binh Chanh",
			Distance:      "19.5 km",
			EstimatedTime: "1 hour 45 minutes",
		},
		{
			StartLocation: "Nguyen Van Cu",
			EndLocation:   "Can Gio Ferry Station",
			Distance:      "54 km",
			EstimatedTime: "3 hours 30 minutes",
		},
		{
			StartLocation: "Phu Lam",
			EndLocation:   "Nha Be",
			Distance:      "25.5 km",
			EstimatedTime: "2 hours",
		},
	}

	// Insert the backofficers into the database
	for _, b := range routes {
		seed.store.CreateRoute(context.Background(), b)
	}

	fmt.Println("Vehicles seed completed.")

	return nil
}
