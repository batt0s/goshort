package controllers

import (
	"encoding/json"
	"net/http"
)

// Function for Sending Response
func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// Request body
type RequestBody struct {
	URL      string `json:"url"`
	Author   string `json:"author"`
	IsCustom bool   `json:"is_custom"`
	Custom   string `json:"custom"`
}
