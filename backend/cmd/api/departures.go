package main

import (
	"net/http"
	"os"
	"fmt"
	"net/url"
	"log"
	"encoding/json"

	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
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
	err := godotenv.Load("../../.env")
	if err != nil{
		log.Println(".env file couldn't be found")
	}
	name:= chi.URLParam(r, "name")
	api_key := os.Getenv("API_KEY") 


	//retrieve stop name	
	if name == ""{
		http.Error(w, "missing stop name", http.StatusBadRequest)
		return
	}

	req := fmt.Sprintf(
		"https://realtime-api.trafiklab.se/v1/stops/name/%s?key=%s",
		url.PathEscape(name),
		api_key,
	) 

	resp, err := http.Get(req)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	error := json.NewDecoder(resp.Body).Decode(&result)
	if error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stop_groups := result["stop_groups"].([]interface{})
	firstStop := stop_groups[0].(map[string]interface{})
	stopID := firstStop["id"].(string)

	//retrieve data

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


	//station departures
	
	var dep_array []Stop
	departures := result1["departures"].([]interface{})
	for i := 0; i<len(departures);i++{
		dep := departures[i].(map[string]interface{
			     
		})
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
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(dep_array)
}
