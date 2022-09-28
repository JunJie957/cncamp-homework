package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	EevVersion = "VERSION"
)

func init() {
	// set os version
	setVersion()
}

func main() {
	http.HandleFunc("/healthz", healthHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// write request header info
	for k, v := range r.Header {
		for _, v2 := range v {
			w.Header().Set(k, v2)
		}
	}

	// write OS VERSION
	w.Header().Set(EevVersion, os.Getenv(EevVersion))

	// write something ...
	_, err := io.WriteString(w, "hello kubernetes ^_^")
	if err != nil {
		log.Fatal(err)
	}

	// log
	log.Println("[INFO] client ip : ", r.Host, ", status code = ", http.StatusOK)
}

func setVersion() {
	err := os.Setenv("VERSION", "1.0.0")
	if err != nil {
		log.Fatal(err)
	}
}
