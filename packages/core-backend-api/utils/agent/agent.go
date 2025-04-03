package agent

import "context"

type Role string
type Action string

const (
	// Roles detailing the type of users in the system
	AIRole     Role = "ai"
	UserRole   Role = "user"
	SystemRole Role = "system"
)

const (
	// Actions detailing the type of actions that can be performed by any of the roles
	ToolCallAction     Action = "toolCall"
	ToolResponseAction Action = "toolResponse"
	TextAction         Action = "text"
)

type HistoryContent struct {
	Action
	Content string // This should be a marschalled json of the action's content
}

type History struct {
	Role
	Content []HistoryContent
}

type StreamingFunc func(ctx context.Context, chunks []byte)

type Tool[Parameter any] interface {
	Name() string
	Description() string
	Actions() []ToolAction[Parameter]
}

type ToolAction[P any] interface {
	Name() string
	Description() string
	Parameters() P
	Call(ctx context.Context, args map[string]any) (response map[string]any, err error)
}

type Agent[Parameter any] interface {
	SetTools(tools []Tool[Parameter])
	SetTemperature(temp float32)
	RunStream(ctx context.Context, prompt string, streamingFunc StreamingFunc) string
}
