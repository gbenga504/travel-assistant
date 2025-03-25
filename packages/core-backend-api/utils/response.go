package utils

func ToErrorResponse(name string, message string) map[string]any {
	return map[string]any{
		"success": false,
		"data": map[string]any{
			"name":    name,
			"message": message,
		},
	}
}

func ToSuccessResponse(data interface{}) map[string]any {
	return map[string]any{
		"success": true,
		"data":    data,
	}
}
