package main

import "github.com/gorilla/mux"

type AuditLoggingService struct {
	router  *mux.Router
	data    AuditLoggingDataService
	metrics ServiceMetrics
}

func NewAuditLoggingService() *AuditLoggingService {
	s := &AuditLoggingService{
		router:  mux.NewRouter(),
		data:    NewMongoDataService(),
		metrics: NewServiceMetrics(),
	}

	s.registerServiceEndpoints()

	return s
}

func init() {
	initMetrics()
}
