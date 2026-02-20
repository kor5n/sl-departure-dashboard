package db

import (
	//"fmt"
	"encoding/json"
	"os"
)

type db struct{
	Stop string `json:stop`
	Routes []string `json:routes`
	Time string `json:time`	
}

type EmptyStruct struct{}

/*func fileExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}*/

func ReadDB()([]db, error){
	data, err :=os.ReadFile("db.json")
	if err != nil{
		return nil, err
	}

	var records []db

	err = json.Unmarshal(data, &records)
	if err != nil{
		return nil, err
	}

	return records, err
}

func WriteToDB(stop string, routes []string, time string){
	newObject := db{
		Stop: stop,
		Routes: routes,
		Time: time,
	}

	records, err := ReadDB()
	if err != nil{
		panic(err)
	}

	records = append(records, newObject)

	updatedData, err := json.MarshalIndent(records, "", " ")
	if err != nil{
		panic(err)
	}

	os.WriteFile("db.json", updatedData, 0644)
}

func DeleteFromDB(index int){
	records, err := ReadDB()
	if err != nil{
		panic(err)
	}

	if index < 0 || index >= len(records) {
		return
	}

	slice := records[:]
	slice = append(slice[:index], slice[index+1:]...)
	copy(records[:], slice)

	updatedData, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("db.json", updatedData, 0664)
	if err != nil {
		panic(err)
	}
}

func Filter(index int)(db, error){
	records, err := ReadDB()
	return records[index], err
}