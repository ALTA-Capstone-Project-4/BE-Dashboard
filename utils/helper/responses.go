package helper

func FailedResponseHelper(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "failed",
		"message": msg,
	}
}

func SuccessResponseHelper(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}

func SuccessDataResponseHelper(msg, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}
