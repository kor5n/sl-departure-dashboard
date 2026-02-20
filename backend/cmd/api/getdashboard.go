package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"backend/internal/db"
	"net/http"
	"strconv"
)

func (api *api) GetDashboard(w http.ResponseWriter, r *http.Request){
	index := chi.URLParam(r,"index")
	idx, err := strconv.Atoi(index)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dboard, err := db.Filter(idx)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dboard)
}
