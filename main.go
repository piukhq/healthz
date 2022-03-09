package main

import (
	"io"
	"log"
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "healthy\n")
}

func main() {
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":6502", nil)
	if err != nil {
		log.Fatal(err)
	}
}
