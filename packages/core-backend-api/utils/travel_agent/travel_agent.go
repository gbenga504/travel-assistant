package travelagent

import (
	"github.com/gbenga504/travel-assistant/utils/agent"
	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	"github.com/gbenga504/travel-assistant/utils/travel_agent/tools/flight"
	"github.com/google/generative-ai-go/genai"
)

const GEMINI_MODEL = "gemini-2.0-flash"

func SetupTravelAgent(gc *gemini.GeminiClient) *gemini.GeminiAgent {
	travelAgent := gemini.NewGeminiAgent(gc, GEMINI_MODEL)

	travelAgent.History = []*agent.History{}

	travelAgent.SetTools([]agent.Tool[*genai.Schema]{
		flight.NewFlightTool(),
	})

	travelAgent.Prompt.ObjectiveAndPersona = objectiveAndPersonaPrompt()
	travelAgent.Prompt.Instructions = instructionsPrompt()
	travelAgent.Prompt.Constraints = constraintsPrompt()
	travelAgent.Prompt.Context = contextPrompt()
	travelAgent.Prompt.OutputFormat = outputFormatPrompt()
	travelAgent.Prompt.Examples = examplesPrompt()

	return travelAgent
}
