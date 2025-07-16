package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gs-jha/Job-Scheduler-Overlap-Checker/endpoints/overlap"
)

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	return &http.Server{Addr: "localhost:" + port, Handler: newHandler()}
}

func newHandler() http.Handler {
	mux := http.NewServeMux()

	overlapHandler := overlap.NewHandler()

	mux.HandleFunc("/api/v1/check-overlap", overlapHandler.CheckOverlap)
	return mux
}
