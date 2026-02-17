package main

import (
	"net/http"
)

func (api *api) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}