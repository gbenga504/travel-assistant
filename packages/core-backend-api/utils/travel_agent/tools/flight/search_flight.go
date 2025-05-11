package flight

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"github.com/gbenga504/travel-assistant/utils/transform"
	"github.com/go-playground/validator/v10"
	"github.com/google/generative-ai-go/genai"
	// googleSearchApi "github.com/serpapi/google-search-results-golang"
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
	// airports := airportNames()

	return &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"departure_city": {Type: genai.TypeString, Description: "Departure city"},
			"arrival_city":   {Type: genai.TypeString, Description: "Arrival city"},
			"departure_date": {Type: genai.TypeString, Description: "Departure date (YYYY-MM-DD)"},
			"return_date":    {Type: genai.TypeString, Description: "Return date (YYYY-MM-DD). Optional parameter. NEVER ask the user for this parameter"},
			"adults":         {Type: genai.TypeNumber, Description: "Number of adults. Optional parameter. NEVER ask the user for this parameter"},
			"max_price":      {Type: genai.TypeNumber, Description: "Maximum ticket price. Optional parameter. NEVER ask the user for this parameter"},
			"max_duration":   {Type: genai.TypeNumber, Description: "Maximum flight duration in hours. Optional parameter. NEVER ask the user for this parameter"},
		},
		Required: []string{"departure_city", "arrival_city", "departure_date"},
	}
}

type ValidatedArgs struct {
	Departure_city string `validate:"required"`
	Arrival_city   string `validate:"required"`
	Departure_date string `validate:"required,datetime=2006-01-02"`
	Return_date    string `validate:"omitempty,datetime=2006-01-02"`
	Adults         int
	Max_price      float64
	Max_duration   float64
}

func (s SearchFlight) Call(ctx context.Context, args map[string]any) (response map[string]any, err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validatedArgs := &ValidatedArgs{}
	transform.MapToExportedStruct(args, validatedArgs)

	err = validate.Struct(validatedArgs)

	if err != nil {
		logger.Fatal("Error validating args for search flight", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrValidatorFailed),
			Message:       errors.Message(errors.ErrValidatorFailed),
			OriginalError: err.Error(),
		})
	}

	// SECRET_API_KEY := utils.LookupEnv("SERP_API_KEY")
	// params := searchParams(*validatedArgs)

	// gSearchApi := googleSearchApi.NewGoogleSearch(params, SECRET_API_KEY)
	// results, err := gSearchApi.GetJSON()

	// if err != nil {
	// 	logger.Fatal("Error searching serp google_flights api", logger.ErrorOpt{
	// 		Name:          errors.Name(errors.ErrThirdPartyAPIRequestFailed),
	// 		Message:       errors.Message(errors.ErrThirdPartyAPIRequestFailed),
	// 		OriginalError: err.Error(),
	// 	})

	// 	return map[string]any{"success": false, "data": nil}, nil
	// }

	////////////// DUMMY DATA ///////////////////
	data, _ := os.ReadFile("./utils/travel_agent/tools/flight/search_flight_mock_data.json")

	var results map[string]any
	if err := json.Unmarshal(data, &results); err != nil {
		logger.Fatal("Error unmarschalling dummy data in search", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrJSONParseIssue),
			Message:       errors.Message(errors.ErrJSONParseIssue),
			OriginalError: err.Error(),
		})
	}
	////////////// DUMMY DATA ///////////////////

	data, err = json.Marshal(results)

	if err != nil {
		logger.Fatal("Error when marschalling response from serp google_flights api", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrJSONParseIssue),
			Message:       errors.Message(errors.ErrJSONParseIssue),
			OriginalError: err.Error(),
		})
	}

	return map[string]any{"success": true, "data": string(data)}, nil
}

func searchParams(args ValidatedArgs) map[string]string {
	airports := airportMap()

	params := map[string]string{
		"engine":        "google_flights",
		"currency":      "EUR",
		"hl":            "en",
		"type":          "2", // This is a one-way trip by default
		"departure_id":  airports[args.Departure_city],
		"arrival_id":    airports[args.Arrival_city],
		"outbound_date": args.Departure_date,
	}

	if args.Return_date != "" {
		params["return_date"] = args.Return_date
	}

	if args.Adults != 0 {
		params["adults"] = strconv.Itoa(args.Adults)
	}

	if args.Max_price != 0 {
		params["max_date"] = fmt.Sprintf("%.2f", args.Max_price)
	}

	if args.Max_duration != 0 {
		params["max_duration"] = fmt.Sprintf("%.2f", args.Max_duration)
	}

	// If the departure and return dates are available then we make this
	// a round-trip
	if args.Departure_date != "" && args.Return_date != "" {
		params["type"] = "1"
	}

	return params
}
