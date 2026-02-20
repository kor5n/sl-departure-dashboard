package db

import (
	//"fmt"
	"encoding/json"
	"os"
	"errors"
)

type db struct{
	StopId string `json:stopid`
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

func WriteToDB(stopid string, routes []string, time string)(error){
	newObject := db{
		StopId: stopid,
		Routes: routes,
		Time: time,
	}

	records, err := ReadDB()
	if err != nil{
		return err
	}

	records = append(records, newObject)

	updatedData, err := json.MarshalIndent(records, "", " ")
	if err != nil{
		return err
	}

	os.WriteFile("db.json", updatedData, 0644)
	return nil
}

func DeleteFromDB(index int)(error){
	records, err := ReadDB()
	if err != nil{
		return err
	}

	if index < 0 || index >= len(records) {
		err := errors.New("wrong index")
		return err
	}

	slice := records[:]
	slice = append(slice[:index], slice[index+1:]...)
	copy(records[:], slice)

	updatedData, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("db.json", updatedData, 0664)
	if err != nil {
		return err
	}

	return nil
}

func Filter(index int)(db, error){
	records, err := ReadDB()
	return records[index], err
}