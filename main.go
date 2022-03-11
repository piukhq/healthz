package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "healthy")
}

func teapot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

func main() {
	http.HandleFunc("/healthz", healthz)

	r := mux.NewRouter()
	r.PathPrefix("/healthz").HandlerFunc(healthz)
	r.PathPrefix("/").HandlerFunc(teapot)
	srv := &http.Server{Handler: r, Addr: ":6502"}
	log.Fatal(srv.ListenAndServe())
}
