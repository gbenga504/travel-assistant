package flight

import (
	"encoding/json"
	"os"

	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
)

func airportMap() map[string]string {
	data, err := os.ReadFile("./utils/travel_agent/tools/flight/airport_codes.json")

	if err != nil {
		logger.Fatal("Error reading airport codes file", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrJSONParseIssue),
			Message:       errors.Message(errors.ErrJSONParseIssue),
			OriginalError: err.Error(),
		})
	}

	var result map[string]string
	if err := json.Unmarshal(data, &result); err != nil {
		logger.Fatal("Error unmarshalling airport codes", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrJSONParseIssue),
			Message:       errors.Message(errors.ErrJSONParseIssue),
			OriginalError: err.Error(),
		})
	}

	return result
}

func airportNames() []string {
	var result []string

	for k := range airportMap() {
		result = append(result, k)
	}

	return result
}
