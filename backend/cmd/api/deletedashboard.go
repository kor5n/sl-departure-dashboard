package main

import (
	"net/http"
	"backend/internal/db"
	"github.com/go-chi/chi/v5"
	"strconv"
)

func (api *api)DeleteDasboard(w http.ResponseWriter, r *http.Request){
	index := chi.URLParam(r, "index")
	idx, err := strconv.Atoi(index)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.DeleteFromDB(idx)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Dashboard deleted"))
}