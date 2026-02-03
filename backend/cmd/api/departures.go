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
	ID int `json:"id"`
	Name string `json:"name"`
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

	departures := result1["departures"].([]interface{})

	fmt.Println(departures)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(departures)
}
