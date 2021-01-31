package handler

import (
	"encoding/json"
	"net/http"
)

var indent bool

// NotFound with a 404 not found status code
func NotFound(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 404)
}

// InternalError with a 500 internal server error.
func InternalError(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 500)
}

// ErrorCode writer a json-encoded error message to http.Respaonse
func ErrorCode(w http.ResponseWriter, err error, status int) {
	JSON(w, &Error{Message: err.Error()}, status)
}

// JSON writer a json-encoded error message to http.Respaonse
func JSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	if indent {
		enc.SetIndent("", "	")
	}
	enc.Encode(v)
}
