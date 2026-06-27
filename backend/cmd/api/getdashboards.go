package main

import (
	"backend/internal/db"
	"encoding/json"
	"net/http"
)

func (api *api) GetDashboards (w http.ResponseWriter, r *http.Request){
	dboards, err := db.ReadDB()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dboards)
}