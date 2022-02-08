package utils

import (
	"bytes"
	"encoding/json"
)

func IndentJsonString(jsonString string) (string, error) {
	var prettyJsonStringMessage bytes.Buffer
	if err := json.Indent(&prettyJsonStringMessage, []byte(jsonString), "", "    "); err != nil {
		return "", err
	}
	return prettyJsonStringMessage.String(), nil
}
