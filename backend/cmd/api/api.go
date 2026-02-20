package main

import (
	"net/http"
	"log"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type api struct{
	config config
}

type config struct{
	addr string
}

func (api *api) mount() http.Handler{
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Use(middleware.Timeout(60 * time.Second))

	//r.Route("/v1", func(r chi.Router){
	r.Get("/health", api.HealthCheckHandler)
	r.Get("/departures/{name}", api.Departures)
	r.Get("/dashboard/{index}", api.GetDashboard)
	/* TODOr.Delete("/delete-dashboard/{index}", api.DeleteDasboard)
	r.Patch("/add-dashboard/") a bunch of arguments
	*/
	//})
	return r
}


func (api *api) run(mux http.Handler) error {
	srv := &http.Server{
		Addr: api.config.addr,
		Handler: mux,
		WriteTimeout:  time.Second*30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("server has started at %s", api.config.addr)

	return srv.ListenAndServe()
}
