package errors

type CustomError string

const (
	ErrEnvNotLoaded CustomError = "ErrEnvNotLoaded"

	ErrAIClientNotLoaded CustomError = "ErrAIClientNotLoaded"
	ErrAIParseIssue      CustomError = "ErrAIParseIssue"
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
	default:
		return "Cannot match error"
	}
}
