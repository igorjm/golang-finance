package http

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/igorjm/golang-finance/adapter/http/transaction"

	"github.com/igorjm/golang-finance/adapter/http/actuator"
)

//Init func
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
