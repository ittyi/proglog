package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer(addr string) *http.Server {
	// httpsrv := newhttpServer()
	r := mux.NewRouter()
	// r.HandleFunc("/", httpsrv.)
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
