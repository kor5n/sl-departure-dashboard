package main

import (
	"backend/internal/db"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (api *api) GetLines(w http.ResponseWriter, r *http.Request){
	name := chi.URLParam(r, "name")
	if name == ""{
		http.Error(w, "missing stop name", http.StatusBadRequest)
		return
	}

	log.Printf("Searching... " + name)
	lines, err := db.LinesByName(name)
	if err != nil{
		if err == db.ErrorStopNotFound{
			http.Error(w, "Stop not found", http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lines)
}