package main

import (
	"encoding/json"
	"github.com/adaptant-labs/opa-audit-logging/middleware"
	"github.com/gorilla/mux"
	opalogs "github.com/open-policy-agent/opa/plugins/logs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func (s *AuditLoggingService) decisionLogHandler(w http.ResponseWriter, r *http.Request) {
	var data []opalogs.EventV1

	params := mux.Vars(r)
	partition := params["partitionName"]

	buf, err  := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}

	log.Infof("Received %d decisions", len(data))

	for i := range data {
		err := s.data.AddDecisionToPartition(partition, data[i])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.metrics.Decisions.Inc()
	}
}

func (s *AuditLoggingService) registerServiceEndpoints() {
	s.router.HandleFunc("/", indexHandler).Methods("GET")
	s.router.Handle("/metrics", promhttp.Handler())
	s.router.HandleFunc("/logs", s.decisionLogHandler).Methods("POST")
	s.router.HandleFunc("/logs/{partitionName}", s.decisionLogHandler).Methods("POST")
	s.router.Use(middleware.TransparentGunzipMiddleware)
}
