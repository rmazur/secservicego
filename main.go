package main

import (
	"log"
	"net/http"
	"os"
	"secservicego/homepage"
	"secservicego/server"
)

var (
	//GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukCertFile = "./certs/local.localhost.cert"
	//GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukKeyFile = "./certs/local.localhost.key"
	//GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
	GcukServiceAddr = "dev.localhost:8080"
)

func main() {
	logger := log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, GcukServiceAddr)
	logger.Println("server starting")
	err := srv.ListenAndServeTLS(GcukCertFile, GcukKeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
