package db

import (
	"encoding/json"
	"os"
	"errors"
	"fmt"
)

type db struct{
	Name string `json:"name"`
	StopId string `json:"stopid"`
	Routes []string `json:"routes"`
	Times []string `json:"times"`	
}

type EmptyStruct struct{}

var DB_PATH string = "internal/db/db.json"

func DBExists() (bool, error) {
	_, err := os.Stat(DB_PATH)

	if os.IsNotExist(err) {
		err := os.WriteFile(DB_PATH, []byte("[]"), 0664)
		if err != nil {
			fmt.Printf("Unable to create db: %v\n", err)
			return false, err
		}
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func ReadDB()([]db, error){
	isdb, err := DBExists()
	if err != nil{
		return nil, err
	}else if isdb != true{
		//database empty
		return nil, nil
		//return nil, errors.New("Database is empty")
	}
	
	data, err :=os.ReadFile(DB_PATH)
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

func WriteToDB(name string,stopid string, routes []string, times []string)(error){
	_, err := DBExists()

	newObject := db{
		Name: name,
		StopId: stopid,
		Routes: routes,
		Times: times,
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

	os.WriteFile(DB_PATH, updatedData, 0644)
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

	records = append(records[:index], records[index+1:]...)

	updatedData, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(DB_PATH, updatedData, 0664)
	if err != nil {
		return err
	}

	return nil
}

func IdxSearch(index int)(db, error){
	records, err := ReadDB()
	if err != nil{
		return db{}, err
	}

	if len(records) <= 0{
		return db{}, errors.New("Database empty")
	}

	if index-1 > len(records){
		return db{}, errors.New("Array index out of range")
	}

	return records[index], err
}

func Filter(value string)(db, error){
	data, err := ReadDB()
	if err != nil{
		return db{}, err
	}

	for _, element := range data {
		if element.StopId == value {
			return element, nil
		}
	}

	return db{}, err
}