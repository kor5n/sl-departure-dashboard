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

func DBExists()(bool,error){
    _, err := os.Stat("db.json")
	if !os.IsNotExist(err) == false{
		err := os.WriteFile("db.json", []byte("{}"), 0664)
		if err != nil {
			fmt.Printf("Uable to create db", err)
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func ReadDB()([]db, error){
	isdb, err := DBExists()
	if isdb != true || err != nil{
		if err != nil{
			return nil, err
		}else if err == nil{
			return nil, errors.New("Database is empty")
		}
	}
	
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

func IdxSearch(index int)(db, error){
	records, err := ReadDB()
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