package seed

import (
	"context"
	"fmt"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
)

func (seed *Seed) runRouteSeed() error {

	
	routes := []db.CreateRouteParams{ 
	{
		StartLocation: "Ben Thanh Market",
		EndLocation: "Cho Lon Bus Station",
		Distance: "6.5 km",
		EstimatedTime: "30 minutes",
	},
	{
		StartLocation: "Ben Thanh Market",
		EndLocation: "An Suong Bus Station",
		Distance: "15.5 km",
		EstimatedTime: "1 hour",
	},
	{
		StartLocation: "Ben Thanh Market",
		EndLocation: "Tan Binh Bus Station",
		Distance: "7.3 km",
		EstimatedTime: "40 minutes",
	},
	{
		StartLocation: "Ben Thanh Market",
		EndLocation: "Nguyen Van Cu Bus Station",
		Distance: "5.2 km",
		EstimatedTime: "25 minutes",
	},
	{
		StartLocation: "Cho Lon Bus Station",
		EndLocation: "Phu Lam Bus Station",
		Distance: "12.5 km",
		EstimatedTime: "1 hour 15 minutes",
	},
	{
		StartLocation: "Cho Lon Bus Station",
		EndLocation: "Hiep Thanh Bus Station",
		Distance: "15.5 km",
		EstimatedTime: "1 hour 30 minutes",
	},
	{
		StartLocation: "An Suong Bus Station",
		EndLocation: "Hoc Mon Bus Station",
		Distance: "25 km",
		EstimatedTime: "2 hours",
	},
	{
		StartLocation: "Tan Binh Bus Station",
		EndLocation: "Binh Chanh Bus Station",
		Distance: "19.5 km",
		EstimatedTime: "1 hour 45 minutes",
	},
	{
		StartLocation: "Nguyen Van Cu Bus Station",
		EndLocation: "Can Gio Ferry Station",
		Distance: "54 km",
		EstimatedTime: "3 hours 30 minutes",
	},
	{
		StartLocation: "Phu Lam Bus Station",
		EndLocation: "Nha Be Bus Station",
		Distance: "25.5 km",
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
