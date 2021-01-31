package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/weekndCN/rw-cert/core"
)

// New a http handler server
func New(info *core.CertInfo) Server {
	return Server{
		CertInfo: info,
	}
}

// Server server Info
type Server struct {
	CertInfo *core.CertInfo
}

// Handler all methods
func (s *Server) Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/list", handleList(*s.CertInfo)).Methods(http.MethodGet)
	r.HandleFunc("/find", handleFind()).Methods(http.MethodGet)
	return r
}
