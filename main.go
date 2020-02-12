package main

import (
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"strconv"
)

var (
	servicePort int
	mongoServer string
)

func main() {
	app := &cli.App{
		Name: "audit-logging-service",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Value:       4123,
				EnvVars:     []string{"AUDIT_LOGGING_SERVICE_PORT"},
				Destination: &servicePort,
			},
			&cli.StringFlag{
				Name:        "mongo-server",
				Value:       "localhost:27017",
				EnvVars:     []string{"AUDIT_LOGGING_SERVICE_MONGO_URI"},
				Destination: &mongoServer,
			},
		},
		Action: func(context *cli.Context) error {
			host := ":" + strconv.Itoa(servicePort)
			srv := NewAuditLoggingService()
			loggedRouter := handlers.LoggingHandler(os.Stdout, srv.router)
			log.Info("Launching new Audit Logging Service on ", host)
			return http.ListenAndServe(host, handlers.RecoveryHandler()(loggedRouter))
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
