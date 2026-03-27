package main

import (
	"net/http"
	"log"
	"time"
	"github.com/go-chi/cors"
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

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router){
		r.Get("/health", api.HealthCheckHandler)
		r.Get("/departures/{id}", api.Departures)
		r.Get("/dashboard/{index}", api.GetDashboard)
		r.Delete("/delete-dashboard/{index}", api.DeleteDasboard)
		r.Post("/add-dashboard/", api.AddDashboard)
		r.Get("/stop-id/{name}", api.GetStopID)
		r.Get("/search-stop/{query}", api.Liststops)
	})
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
