package main

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
	"log"
)

type Stop struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Canceled bool `json:"canceled"`
	Route string `json:"route"`
	Direction string `json:"direction"`
	Mode string `json:"mode"`
	Departure string `json:"departure"`
	Stop string `json:"stop"`
	Alerts []interface{} `json:"alerts"`
}
func (api *api) Departures(w http.ResponseWriter, r *http.Request){

	//get api key from .env
	err := godotenv.Load("../../.env")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	api_key := os.Getenv("API_KEY")
	stopID := chi.URLParam(r, "id")

	//retrieve departures
	req1 := fmt.Sprintf(
		"https://realtime-api.trafiklab.se/v1/departures/%s?key=%s",
		stopID,
		api_key,
	)

	resp1,err := http.Get(req1)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp1.Body.Close()
	
	var result1 map[string]interface{}
	error1 := json.NewDecoder(resp1.Body).Decode(&result1)
	if error1 != nil {
		http.Error(w, error1.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(result1)
	stops := result1["stops"].([]interface{})
	stop := stops[0].(map[string]interface{})
	name := stop["name"].(string)
	//station departures
	
	var dep_array []Stop
	departures := result1["departures"].([]interface{})
	for i := 0; i<len(departures);i++{
		dep := departures[i].(map[string]interface{})
		canceled := dep["canceled"].(bool)
		routepath := dep["route"].(map[string]interface{})
		route := routepath["designation"].(string)
		direction := routepath["direction"].(string)
		mode := routepath["transport_mode"].(string)
		realtime := dep["realtime"].(string)
		plat := dep["realtime_platform"].(map[string]interface{})
		stop := plat["designation"].(string)
		alerts := dep["alerts"].([]interface{})

		station := Stop{
			ID: stopID,
			Name: name,
			Canceled: canceled,
			Route: route,
			Direction: direction,
			Mode: mode,
			Departure: realtime,
			Stop: stop,
			Alerts: alerts,
		}
		dep_array = append(dep_array, station)
	}
	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp1.StatusCode)
	json.NewEncoder(w).Encode(dep_array)
}
