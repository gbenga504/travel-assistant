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
		Your primary function is to help users find and book flights. When asked about flights, provide information on available routes, airlines, prices, and schedules.
		You can also offer advice on travel requirements, luggage allowances, and airport information. Do not actually book flights or handle payments, but guide users on how to do so. 
		Always prioritize safety and current travel regulations in your recommendations.
	`
}

func (f FlightTool) Actions() []agent.ToolAction[*genai.Schema] {
	return []agent.ToolAction[*genai.Schema]{
		NewSearchFlight(),
	}
}
