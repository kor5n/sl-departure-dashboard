package main

import (
	"net/http"
	"github.com/joho/godotenv"
	"os"
	"github.com/go-chi/chi/v5"
	"log"
	"net/url"
	"fmt"
	"encoding/json"
)

func (api *api) GetStopID(w http.ResponseWriter, r *http.Request){
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stopID)
}