package main

import (
	"encoding/json"
	"net/http"
	"backend/internal/db"
)

type Request struct {
    StopId   string   `json:"stopid"`
    Routes []string `json:"routes"`
    Time   string   `json:"time"`
}


func (api *api)AddDashboard(w http.ResponseWriter, r *http.Request){
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.WriteToDB(req.StopId, req.Routes, req.Time)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Added new dashboard"))
}