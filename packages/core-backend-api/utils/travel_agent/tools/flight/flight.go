package flight

import (
	"github.com/gbenga504/travel-assistant/utils/agent"
	llmcontext "github.com/gbenga504/travel-assistant/utils/llm_context"
	"github.com/google/generative-ai-go/genai"
)

type FlightTool struct {
	llmContext *llmcontext.LLMContext
}

var _ agent.Tool[*genai.Schema] = (*FlightTool)(nil)

func NewFlightTool(llmContext *llmcontext.LLMContext) FlightTool {
	return FlightTool{llmContext: llmContext}
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
		NewSearchFlight(f.llmContext),
	}
}
