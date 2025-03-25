package errors

type CustomError string

const (
	ErrEnvNotLoaded CustomError = "ErrEnvNotLoaded"

	ErrAIClientNotLoaded CustomError = "ErrAIClientNotLoaded"
	ErrAIParseIssue      CustomError = "ErrAIParseIssue"
	ErrServerClosed      CustomError = "ErrServerClosed"
	ErrDatabaseIssue     CustomError = "ErrDatabaseIssue"
	ErrJSONParseIssue    CustomError = "ErrJSONParseIssue"
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
	case ErrJSONParseIssue:
		return "Cannot parse JSON"
	default:
		return "Cannot match error"
	}
}
