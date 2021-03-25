package xsql

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsQuery = promauto.NewCounter(prometheus.CounterOpts{
		Name: "query_request",
		Help: "read count",
	})

	opsExec = promauto.NewCounter(prometheus.CounterOpts{
		Name: "exec_request",
		Help: "write count",
	})

	opsTx = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tx_request",
		Help: "tx count",
	})
)
