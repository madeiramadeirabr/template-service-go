package healthCheck

func GetStatus() map[string]interface{} {
	return map[string]interface{}{
		"status": "OK",
	}
}
