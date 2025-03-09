package gemini

import (
	"encoding/json"

	"github.com/gbenga504/travel-assistant/utils/agent"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"github.com/google/generative-ai-go/genai"
)

func (ga *GeminiAgent) addToHistory(mc *ModifiedGenAIContent) {
	history := agent.History{
		Role: mc.Role,
	}

	history.Content = genAIPartsToHistoryContents(mc.Parts)
	ga.History = append(ga.History, &history)

	// We have to notify about the history change if the ListenAndNotifyHistoryChange
	// method is available
	if ga.ListenAndNotifyHistoryChange != nil {
		ga.ListenAndNotifyHistoryChange(history)
	}
}

func (ga *GeminiAgent) historyToGenAIHistory() []*genai.Content {
	var result []*genai.Content

	for _, h := range ga.History {
		role := "model"

		if h.Role == agent.SystemRole {
			role = "user"
		}

		content := genai.Content{
			Role:  role,
			Parts: historyContentsToGenAIParts(h.Content),
		}

		result = append(result, &content)
	}

	return result
}

func genAIPartsToHistoryContents(parts []genai.Part) []agent.HistoryContent {
	var historyContents []agent.HistoryContent

	for _, part := range parts {
		switch v := part.(type) {
		case genai.Text:
			historyContents = append(historyContents, agent.HistoryContent{Action: agent.TextAction, Content: string(v)})

		case genai.FunctionCall:
			{
				js, err := json.Marshal(map[string]any{
					"name": v.Name,
					"args": v.Args,
				})

				if err != nil {
					logger.Error("Failed to marshal the content of a function call", logger.ErrorOpt{
						Name:          errors.Name(errors.ErrAIParseIssue),
						Message:       errors.Message(errors.ErrAIParseIssue),
						OriginalError: err.Error(),
					})
				}

				historyContents = append(historyContents, agent.HistoryContent{Action: agent.ToolCallAction, Content: string(js)})
			}

		case genai.FunctionResponse:
			{
				js, err := json.Marshal(map[string]any{
					"name":     v.Name,
					"response": v.Response,
				})

				if err != nil {
					logger.Error("Failed to marshal the content of a function response", logger.ErrorOpt{
						Name:          errors.Name(errors.ErrAIParseIssue),
						Message:       errors.Message(errors.ErrAIParseIssue),
						OriginalError: err.Error(),
					})
				}

				historyContents = append(historyContents, agent.HistoryContent{Action: agent.ToolResponseAction, Content: string(js)})
			}
		}
	}

	return historyContents
}

func historyContentsToGenAIParts(hc []agent.HistoryContent) []genai.Part {
	var parts []genai.Part

	for _, c := range hc {
		switch c.Action {
		case agent.TextAction:
			parts = append(parts, genai.Text(c.Content))

		case agent.ToolCallAction:
			{
				var funcCall struct {
					Name string
					Args map[string]any
				}

				if err := json.Unmarshal([]byte(c.Content), &funcCall); err != nil {
					logger.Error("Failed to unmarschal the content of a tool call", logger.ErrorOpt{
						Name:          errors.Name(errors.ErrAIParseIssue),
						Message:       errors.Message(errors.ErrAIParseIssue),
						OriginalError: err.Error(),
					})
				}

				parts = append(parts, genai.FunctionCall{
					Name: funcCall.Name,
					Args: funcCall.Args,
				})
			}

		case agent.ToolResponseAction:
			{
				var funcResp struct {
					Name     string
					Response map[string]any
				}

				if err := json.Unmarshal([]byte(c.Content), &funcResp); err != nil {
					logger.Error("Failed to unmarschal the content of a tool response", logger.ErrorOpt{
						Name:          errors.Name(errors.ErrAIParseIssue),
						Message:       errors.Message(errors.ErrAIParseIssue),
						OriginalError: err.Error(),
					})
				}

				parts = append(parts, genai.FunctionResponse{
					Name:     funcResp.Name,
					Response: funcResp.Response,
				})
			}
		}
	}

	return parts
}
