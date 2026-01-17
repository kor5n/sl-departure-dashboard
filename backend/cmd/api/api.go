package main

import (
	"net/http"
	"log"
	"time"
)

type api struct{
	config config
}

type config struct{
	addr string
}

func (api *api) mount() *http.
ServeMux{
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health",api.healthCheckHandler)
	return mux
}


func (api *api) run(mux *http.ServeMux) error {
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
