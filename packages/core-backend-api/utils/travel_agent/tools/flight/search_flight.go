package flight

import (
	"context"

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
		Search for flights based on user input. Provide options including routes, airlines, prices, and schedules. Do not book flights.
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
	return map[string]any{"flights": []map[string]any{
		{"flightId": "F1", "price": 500, "airline": "Airline A"},
		{"flightId": "F2", "price": 450, "airline": "Airline B"},
		{"flightId": "F3", "price": 550, "airline": "Airline C"},
	}}, nil
}
