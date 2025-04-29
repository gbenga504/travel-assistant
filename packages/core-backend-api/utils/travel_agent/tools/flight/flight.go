package flight

import (
	"github.com/gbenga504/travel-assistant/utils/agent"
	"github.com/google/generative-ai-go/genai"
)

type FlightTool struct{}

var _ agent.Tool[*genai.Schema] = (*FlightTool)(nil)

func NewFlightTool() FlightTool {
	return FlightTool{}
}

func (f FlightTool) Name() string {
	return "flight_tool"
}

func (f FlightTool) Description() string {
	return `
		Your primary function is to assist users with their flight-related queries.
	`
}

func (f FlightTool) Actions() []agent.ToolAction[*genai.Schema] {
	return []agent.ToolAction[*genai.Schema]{
		NewSearchFlight(),
	}
}
