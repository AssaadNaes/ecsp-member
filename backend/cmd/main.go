package main

import (
	"net/http"
	"os"
	"team_members/internal"
	"time"

	"github.com/labstack/gommon/log"
)

func main() {
	api := internal.NewApi()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /member", api.HandleGetMember)

	port := os.Getenv("MEMBER_PORT")
	if port == "" {
		port = "8000"
	}

	srv := http.Server{
		Handler:           mux,
		Addr:              ":" + port,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Info("Starting server on port: ", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
