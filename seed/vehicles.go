package seed

import (
	"context"
	"fmt"
	"log"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
)

func (seed *Seed) runVehicleSeed() error {

	vehicles := []db.CreateVehicleParams{
		{
			MakeBy:          "Ford",
			Model:           "F-150",
			Capacity:        "3 tons",
			FuelConsumption: "11.2 L/100km",
		},
		{
			MakeBy:          "Chevrolet",
			Model:           "Silverado",
			Capacity:        "3 tons",
			FuelConsumption: "12.8 L/100km",
		},
		{
			MakeBy:          "Ram",
			Model:           "1500",
			Capacity:        "3 tons",
			FuelConsumption: "12.1 L/100km",
		},
		{
			MakeBy:          "GMC",
			Model:           "Sierra",
			Capacity:        "3 tons",
			FuelConsumption: "13.5 L/100km",
		},
		{
			MakeBy:          "Toyota",
			Model:           "Tacoma",
			Capacity:        "1.5 tons",
			FuelConsumption: "10.4 L/100km",
		},
		{
			MakeBy:          "Nissan",
			Model:           "Frontier",
			Capacity:        "1.5 tons",
			FuelConsumption: "11.8 L/100km",
		},
		{
			MakeBy:          "Isuzu",
			Model:           "D-Max",
			Capacity:        "2.5 tons",
			FuelConsumption: "9.5 L/100km",
		},
		{
			MakeBy:          "Mitsubishi",
			Model:           "L200",
			Capacity:        "2.5 tons",
			FuelConsumption: "8.9 L/100km",
		},
		{
			MakeBy:          "Mercedes-Benz",
			Model:           "Actros",
			Capacity:        "40 tons",
			FuelConsumption: "32 L/100km",
		},
		{
			MakeBy:          "Volvo",
			Model:           "FH16",
			Capacity:        "40 tons",
			FuelConsumption: "30 L/100km",
		},
	}

	for _, b := range vehicles {
		ctx := context.Background()
		vehicle, _ := seed.store.CreateVehicle(ctx, b)
		_, err := seed.store.CreateVehicleStatus(ctx, int32(vehicle.ID))
		if err != nil {
			log.Fatal("error: ", err)
		}

	}

	fmt.Println("Vehicles seed completed.")

	return nil
}
