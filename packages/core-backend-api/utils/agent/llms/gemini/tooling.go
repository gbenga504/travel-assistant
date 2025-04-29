package gemini

import (
	"context"
	"time"

	"github.com/gbenga504/travel-assistant/utils/agent"
	"github.com/google/generative-ai-go/genai"
)

type ToolActionResponse struct {
	Name     string         // This is the name of the tool action
	Response map[string]any // This is the response of the tool action
}

func handleFunctionCall(ctx context.Context, funcCalls []genai.FunctionCall, tools []agent.Tool[*genai.Schema]) []genai.Part {
	var result []genai.Part
	c := make(chan ToolActionResponse)

	for _, fc := range funcCalls {
		toolAction := retrieveToolAction(fc.Name, tools)
		runToolAction(toolAction, ctx, fc.Args, c)
	}

	tr := aggregateToolResponses(c, len(funcCalls))

	for _, r := range tr {
		result = append(result, genai.FunctionResponse{
			Name:     r.Name,
			Response: r.Response,
		})
	}

	return result
}

func runToolAction(t agent.ToolAction[*genai.Schema], ctx context.Context, args map[string]any, writerChan chan<- ToolActionResponse) {
	go func() {
		resp, err := t.Call(ctx, args)

		result := ToolActionResponse{
			Name: t.Name(),
		}

		if err != nil {
			result.Response = map[string]any{
				"success": "false",
				"error": map[string]any{
					"name": err.Error(),
				},
			}
		} else {
			result.Response = resp
		}

		writerChan <- result
	}()
}

func retrieveToolAction(functionCallName string, tools []agent.Tool[*genai.Schema]) agent.ToolAction[*genai.Schema] {
	var toolAction agent.ToolAction[*genai.Schema]

	for _, tool := range tools {
		for _, action := range tool.Actions() {
			if action.Name() == functionCallName {
				toolAction = action

				break
			}
		}

		if toolAction != nil {
			break
		}
	}

	return toolAction
}

func aggregateToolResponses(reader <-chan ToolActionResponse, numToolActionCalls int) []ToolActionResponse {
	var result []ToolActionResponse

	for {
		var done bool

		select {
		case r := <-reader:
			result = append(result, r)

			// We check if we have completed aggregating all tool actions and then abort if true
			if len(result) == numToolActionCalls {
				done = true
			}

		// If we don't get any response after 5 seconds then we need to abort
		case <-time.After(5 * time.Second):
			done = true
		}

		if done {
			break
		}
	}

	return result
}
