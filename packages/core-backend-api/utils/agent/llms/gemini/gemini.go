package gemini

import (
	"context"
	"strings"

	"github.com/gbenga504/travel-assistant/utils/agent"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"github.com/google/generative-ai-go/genai"
)

type GeminiAgent struct {
	tools                        []agent.Tool[*genai.Schema]
	model                        *genai.GenerativeModel
	chatSession                  *genai.ChatSession
	History                      []*agent.History
	ListenAndNotifyHistoryChange (func(history agent.History))
	agent.Prompt
}

// This is an adaptation of genai.Content
// We use this because we want to replace the Role field with our Role enum
type ModifiedGenAIContent struct {
	agent.Role
	Parts []genai.Part
}

var _ agent.Agent[*genai.Schema] = (*GeminiAgent)(nil)

func NewGeminiAgent(gc *GeminiClient, modelName string) *GeminiAgent {
	model := gc.client.GenerativeModel(modelName)

	// We set the base config on the model, so in case it is not passed
	// we can default to these values
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(4000)
	model.ResponseMIMEType = "text/plain"

	return &GeminiAgent{
		model:       model,
		chatSession: model.StartChat(),
	}
}

func (ga *GeminiAgent) SetTemperature(temp float32) {
	ga.model.SetTemperature(temp)
}

func (ga *GeminiAgent) SetTopK(topK int32) {
	ga.model.SetTopK(topK)
}

func (ga *GeminiAgent) SetTopP(topP float32) {
	ga.model.SetTopP(topP)
}

func (ga *GeminiAgent) SetMaxOutputTokens(maxOutputTokens int32) {
	ga.model.SetMaxOutputTokens(maxOutputTokens)
}

func (ga *GeminiAgent) SetResponseMIMEType(responseMIMEType string) {
	ga.model.ResponseMIMEType = responseMIMEType
}

func (ga *GeminiAgent) SetTools(tools []agent.Tool[*genai.Schema]) {
	var genaiTools []*genai.Tool

	for _, tool := range tools {
		genaiTools = append(genaiTools, &genai.Tool{FunctionDeclarations: actionsToFunctionDeclarations((tool).Actions())})
	}

	ga.tools = tools
	ga.model.Tools = genaiTools
}

func (ga *GeminiAgent) RunStream(ctx context.Context, userPrompt string, streamingFunc agent.StreamingFunc) {
	// Set the system instructions
	ga.model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(ga.Prompt.Stitch())},
	}

	ga.chatSession.History = ga.historyToGenAIHistory()

	prompt := genai.Text(userPrompt)

	// Save the initial text from the user in the history
	ga.addToHistory(&ModifiedGenAIContent{
		Role:  agent.UserRole,
		Parts: []genai.Part{prompt},
	})

	messageStream := ga.chatSession.SendMessageStream(ctx, prompt)
	ga.processStream(ctx, prompt, messageStream, streamingFunc)
}

func (ga *GeminiAgent) processStream(ctx context.Context, userPrompt genai.Text, ms *genai.GenerateContentResponseIterator, sf agent.StreamingFunc) {
	resp, err := ms.Next()

	for resp != nil {
		// TODO: Evaluate if it makes sense to do this or just call the streaming Func with the error
		if err != nil {
			logger.Error("Error processing stream", logger.ErrorOpt{
				Name:          errors.Name(errors.ErrAIParseIssue),
				Message:       errors.Message(errors.ErrAIParseIssue),
				OriginalError: err.Error(),
			})
		}

		// It is important to stream text responses as early as possible
		streamTextResponse(ctx, resp.Candidates[0].Content.Parts, sf)

		resp, err = ms.Next()
		mergedResp := ms.MergedResponse()

		// When the stream is done, it is important to peform some actions
		// E.g we need to update history, run any function calls and pass the function response to the model
		if resp == nil {
			// Update History
			ga.addToHistory(&ModifiedGenAIContent{
				Role:  agent.AIRole,
				Parts: mergedResp.Candidates[0].Content.Parts,
			})

			// Handle Function Calls
			funcCalls := mergedResp.Candidates[0].FunctionCalls()
			if len(funcCalls) > 0 {
				fcParts := handleFunctionCall(ctx, funcCalls, ga.tools)

				parts := []genai.Part{userPrompt}
				parts = append(parts, fcParts...)

				ga.addToHistory(&ModifiedGenAIContent{
					Role:  agent.SystemRole,
					Parts: parts,
				})

				messageStream := ga.chatSession.SendMessageStream(ctx, parts...)
				resp, err = messageStream.Next()
			}
		}
	}
}

func streamTextResponse(ctx context.Context, parts []genai.Part, sf agent.StreamingFunc) {
	for _, part := range parts {
		if text, ok := part.(genai.Text); ok {
			if strings.TrimSpace(string(text)) == "" {
				continue
			}

			sf(ctx, []byte(text))
		}
	}
}

func actionsToFunctionDeclarations(actions []agent.ToolAction[*genai.Schema]) []*genai.FunctionDeclaration {
	var funcDeclarations []*genai.FunctionDeclaration

	for _, action := range actions {

		funcDeclarations = append(funcDeclarations, &genai.FunctionDeclaration{
			Name:        action.Name(),
			Description: action.Description(),
			Parameters:  action.Parameters(),
		})

	}

	return funcDeclarations
}
