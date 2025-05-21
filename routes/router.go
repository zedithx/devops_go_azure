package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zedithx/devops_practice/handlers"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", handlers.HelloHandler)
	r.Get("/status", handlers.StatusHandler)
	r.Handle("/metrics", promhttp.Handler())

	return r 
}