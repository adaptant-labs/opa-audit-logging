package main

import "github.com/gorilla/mux"

type AuditLoggingService struct {
	router  *mux.Router
	metrics ServiceMetrics
}

func NewAuditLoggingService() *AuditLoggingService {
	s := &AuditLoggingService{
		router:  mux.NewRouter(),
		metrics: NewServiceMetrics(),
	}

	s.registerServiceEndpoints()

	return s
}

func init() {
	initMetrics()
}