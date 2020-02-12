package main

import opalogs "github.com/open-policy-agent/opa/plugins/logs"

type AuditLoggingDataService interface {
	AddDecision(decision opalogs.EventV1) error
	AddDecisionToPartition(partition string, decision opalogs.EventV1) error
}
