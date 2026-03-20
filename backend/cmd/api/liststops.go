package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func contains[T comparable](arr []T, target T) bool {
    for _, v := range arr {
        if v == target {
            return true
        }
    }
    return false
}

func (api *api) Liststops(w http.ResponseWriter, r *http.Request){
	err := godotenv.Load("../../.env")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	api_key := os.Getenv("API_KEY")

	query := chi.URLParam(r,"query")

	req := fmt.Sprintf(
		"https://realtime-api.trafiklab.se/v1/stops/name/%s?key=%s",
		query,
		api_key,
	)

	resp, err := http.Get(req)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	
	var result map[string]interface{}
	error1 := json.NewDecoder(resp.Body).Decode(&result)
	if error1 != nil {
		http.Error(w, error1.Error(), http.StatusInternalServerError)
		return
	}

	var stop_names []string
	stops := result["stop_groups"].([]interface{})

	for i := 0; i<len(stops); i++{
		stop := stops[i].(map[string]interface{})
		if contains(stop_names, stop["name"].(string)) == false{
			stop_names = append(stop_names, stop["name"].(string))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if len(stop_names) == 0{
		stop_names = append(stop_names, "Couldn't find stops")
	}

	json.NewEncoder(w).Encode(stop_names)

}