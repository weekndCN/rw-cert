package handler

import (
	"encoding/json"
	"net/http"

	"github.com/weekndCN/rw-cert/core"
)

// HandleFind find a host
func handleFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// HandleList list all hosts
func handleList(info core.CertInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(info)
		if err != nil {
			InternalError(w, err)
		}
		JSON(w, string(b), 200)
	}
}
