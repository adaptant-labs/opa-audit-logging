package main

import (
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	log.Info("Launching audit logging service..")
	srv := NewAuditLoggingService()
	loggedRouter := handlers.LoggingHandler(os.Stdout, srv.router)
	log.Fatal(http.ListenAndServe(":4123", handlers.RecoveryHandler()(loggedRouter)))
}
