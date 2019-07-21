package utils

import (
	"net/http"
	"encoding/json"
)

func SendError(w http.ResponseWriter, msg string) {
	w.Header().Add("Content-Type", "application/json")

	respMap := map[string] interface{} {
		"success": false, 
		"message": msg,
	}

	json.NewEncoder(w).Encode(respMap);
}

func SendSuccess(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")

	respMap := map[string] interface{} {
		"success": true,
	}

	for k, v := range data {
    respMap[k] = v
  }

	json.NewEncoder(w).Encode(respMap);
}