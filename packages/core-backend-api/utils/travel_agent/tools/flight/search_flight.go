package flight

import (
	"context"
	"encoding/json"

	"github.com/google/generative-ai-go/genai"
)

type SearchFlight struct{}

func NewSearchFlight() SearchFlight {
	return SearchFlight{}
}

func (s SearchFlight) Name() string {
	return "search_flights"
}

func (s SearchFlight) Description() string {
	return `
		Search for flights based on user input. Provide options including routes, airlines, prices, and schedules.
	`
}

func (s SearchFlight) Parameters() *genai.Schema {
	return &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"origin":      {Type: genai.TypeString, Description: "Departure city"},
			"destination": {Type: genai.TypeString, Description: "Arrival city"},
			"date":        {Type: genai.TypeString, Description: "Flight date (YYYY-MM-DD)"},
		},
		Required: []string{"origin", "destination", "date"},
	}
}

func (s SearchFlight) Call(ctx context.Context, args map[string]any) (response map[string]any, err error) {
	// parameter := map[string]string{
	// 	"engine":        "google_flights",
	// 	"departure_id":  "PEK",
	// 	"arrival_id":    "AUS",
	// 	"outbound_date": "2025-04-28",
	// 	"return_date":   "2025-05-04",
	// 	"currency":      "USD",
	// 	"hl":            "en",
	// }

	// gSearchApi := googleSearchApi.NewGoogleSearch(parameter, "secret_api_key")
	// results, err := gSearchApi.GetJSON()

	// fmt.Printf("gad the results from %#v and the error is %#v\n", results, err.Error())

	a, err := json.Marshal([]map[string]any{
		{"flightId": "F1", "price": 500, "airline": "Airline A"},
		{"flightId": "F2", "price": 450, "airline": "Airline B"},
		{"flightId": "F3", "price": 550, "airline": "Airline C"},
	})

	return map[string]any{"flights": string(a)}, err
}
