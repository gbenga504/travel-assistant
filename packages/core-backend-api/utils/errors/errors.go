package errors

type CustomError string

const (
	ErrEnvNotLoaded CustomError = "ErrEnvNotLoaded"

	ErrAIClientNotLoaded CustomError = "ErrAIClientNotLoaded"
	ErrAIParseIssue      CustomError = "ErrAIParseIssue"
	ErrServerClosed      CustomError = "ErrServerClosed"
	ErrDatabaseIssue     CustomError = "ErrDatabaseIssue"
)

func Name(err CustomError) string {
	return string(err)
}

func Message(err CustomError) string {
	switch err {
	case ErrEnvNotLoaded:
		return "Env variable is missing"
	case ErrAIClientNotLoaded:
		return "AI client was not loaded"
	case ErrAIParseIssue:
		return "Cannot parse AI related data"
	case ErrServerClosed:
		return "Server was closed"
	case ErrDatabaseIssue:
		return "Database had an issue"
	default:
		return "Cannot match error"
	}
}

func ToErrorResponse(name string, message string) map[string]any {
	return map[string]any{
		"success": false,
		"data": map[string]any{
			"name":    name,
			"message": message,
		},
	}
}
