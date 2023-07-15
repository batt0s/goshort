package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Function for Sending Response
func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// RequestBody : Request body
type RequestBody struct {
	URL    string `json:"url"`
	Custom string `json:"custom"`
}

// The source that helped me to write this helper: https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

func getRequestBody(w http.ResponseWriter, r *http.Request) (*RequestBody, error) {
	// Check if the Content-Type is json
	if r.Header.Get("Content-Type") != "application/json" {
		msg := "content-type is not application/json"
		return nil, &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	var body RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&body)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshallTypeError *json.UnmarshalTypeError
		switch {
		case errors.As(err, &syntaxError):
			msg := "Request body contains badly-formed JSON"
			return nil, &malformedRequest{status: http.StatusBadRequest, msg: msg}
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			return nil, &malformedRequest{status: http.StatusBadRequest, msg: msg}
		case errors.As(err, &unmarshallTypeError):
			msg := fmt.Sprintf("Request body contains invalid value for the %q", unmarshallTypeError.Field)
			return nil, &malformedRequest{status: http.StatusBadRequest, msg: msg}
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			field := strings.TrimPrefix(err.Error(), "json: unknown field")
			msg := fmt.Sprintf("Request body containt unknown field %s", field)
			return nil, &malformedRequest{status: http.StatusBadRequest, msg: msg}
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return nil, &malformedRequest{status: http.StatusBadRequest, msg: msg}
		case err.Error() == "http: request body too large":
			msg := "Request body must not be longer than 1MB"
			return nil, &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}
		default:
			return nil, err
		}
	}
	return &body, nil
}
