package db

import (
	"encoding/json"
	"errors"
	"os"
	"slices"
)

var ErrorStopNotFound = errors.New("No stop found")


type Route struct {
	RouteId string `json:"route_id"`
	Line string `json:"line"`
	Type int `json:"type"`
}

type Stop struct{
	StopName string `json:"stop_name"`
	Routes []Route `json:"routes"` 
}

func LinesByName(stop_name string)([]string, error){
	data, err:= os.ReadFile("internal/db/stop_info.json")
	if err != nil{
		return nil, err
	}

	var stops map[string]Stop

	err1 := json.Unmarshal(data, &stops)
	if err1 != nil{
		return nil, err1
	}

	var foundLines []string
	for _,stop := range stops{
		if stop.StopName == stop_name{
			for _, route := range stop.Routes{
				if !slices.Contains(foundLines, route.Line){
					foundLines = append(foundLines, route.Line)
				}
			}
		} 
	}

	if len(foundLines) == 0{
		return nil, ErrorStopNotFound
	}

	return foundLines, nil
}