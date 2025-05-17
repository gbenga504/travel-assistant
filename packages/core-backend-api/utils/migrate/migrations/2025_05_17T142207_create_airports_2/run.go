package create_airports_2_2025_05_17T142207

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CreateAirports220250517T142207 struct{}

type Airport struct {
	City    string
	State   *string
	Country string
	Code    string
}

func (m *CreateAirports220250517T142207) Up(db *mongo.Database) error {
	// Read the airports json and unmarschal
	data, err := os.ReadFile("./migrations/2025_05_16T221702_create_airports_1/airports.json")

	if err != nil {
		return err
	}

	var airports []Airport
	err = json.Unmarshal(data, &airports)

	if err != nil {
		return err
	}

	groupedAirports := groupAirportsByCities(airports)
	_, err = db.Collection("airports").InsertMany(context.Background(), groupedAirports)

	if err != nil {
		return err
	}

	return nil
}

func (m *CreateAirports220250517T142207) Down(db *mongo.Database) error {
	fmt.Println("Noop")

	return nil
}

type AirportsGroupedByCity struct {
	City    string
	State   *string
	Country string
	Codes   []string
}

func groupAirportsByCities(airports []Airport) []AirportsGroupedByCity {
	var groupedAirportsMap = make(map[string]AirportsGroupedByCity)

	for _, airport := range airports {
		if groupedAirports, exists := groupedAirportsMap[airport.City]; exists {
			// Also need to make sure that the state is the same
			// First we have to derefernce the value of each state before comparing
			var airportState string
			var groupedAirportsState string

			if airport.State != nil {
				airportState = *airport.State
			}

			if groupedAirports.State != nil {
				groupedAirportsState = *groupedAirports.State
			}

			if airportState == groupedAirportsState {
				groupedAirportsMap[airport.City] = AirportsGroupedByCity{
					City:    airport.City,
					State:   airport.State,
					Country: airport.Country,
					Codes:   append(groupedAirports.Codes, airport.Code),
				}

				continue
			}
		}

		groupedAirportsMap[airport.City] = AirportsGroupedByCity{
			City:    airport.City,
			State:   airport.State,
			Country: airport.Country,
			Codes:   []string{airport.Code},
		}
	}

	var result = make([]AirportsGroupedByCity, 0, len(groupedAirportsMap))
	for _, groupedAirports := range groupedAirportsMap {
		result = append(result, groupedAirports)
	}

	return result
}
