package main

import (
	"net/http"
	"os"
	"fmt"
	"net/url"
	"log"
	"encoding/json"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
)

type Stop struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func (api *api) StopID(w http.ResponseWriter, r *http.Request){
	err := godotenv.Load("../../.env")
	if err != nil{
		log.Println(".env file couldn't be found")
	}
	name:= chi.URLParam(r, "name")
	api_key := os.Getenv("API_KEY") 

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

	id, err := strconv.Atoi(stopID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stop := Stop{
		ID: id,
		Name: name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(stop)
}
