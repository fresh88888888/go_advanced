package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CorsConfig struct {
	AllowedMethods  []string
	Allowedheaders  []string
	WithCredentials bool
}

var beforeAllowedMethods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodDelete,
	http.MethodPatch,
	http.MethodPost,
	http.MethodPut,
}

const defaultOrigin = "*"

func main() {
	c := &CorsConfig{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleCorsRequest(c))

	server := &http.Server{
		Addr:           ":8090",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

func handleCorsRequest(c *CorsConfig) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		w.Header().Set("Accesee-Control-Allow-Origin", defaultOrigin)

		body := proxyRequest(http.DefaultClient, r, ctx)
		fmt.Fprintf(w, "%s", body)
	}
}

func proxyRequest(client *http.Client, r *http.Request, ctx context.Context) string {
	// perform request and read body here
	return ""
}
