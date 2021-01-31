package handler

import (
	"fmt"
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
		// if config not set
		if len(info.Info) == 0 {
			NotFound(w, fmt.Errorf("no result founded"))
			return
		}
		JSON(w, info.Info, 200)
	}
}
