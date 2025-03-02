package errors

type CustomError string

const (
	EnvLoadError CustomError = "EnvLoadError"
)

func Name(err CustomError) string {
	return string(err)
}

func Message(err CustomError) string {
	switch err {
	case EnvLoadError:
		return "Env variable is missing"
	default:
		return "Cannot match error"
	}
}
