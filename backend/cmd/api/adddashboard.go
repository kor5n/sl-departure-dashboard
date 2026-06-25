package main

import (
	"backend/internal/db"
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Name string `json:name`
    StopId   string   `json:"stopid"`
    Routes []string `json:"routes"`
    Times   []string   `json:"times"`
}


func (api *api)AddDashboard(w http.ResponseWriter, r *http.Request){
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(req)
	err = db.WriteToDB(req.Name,req.StopId, req.Routes, req.Times)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Added new dashboard"))
}