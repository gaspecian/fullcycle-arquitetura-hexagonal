package handler

import "encoding/json"

func jsonError(msg string) []byte {
	errorJson := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	r, err := json.Marshal(errorJson)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
