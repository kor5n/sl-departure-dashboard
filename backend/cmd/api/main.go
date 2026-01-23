package main

import (
	"log"
)

func main(){
	cfg := config{
		addr: ":8080",
	}

	api := &api{
		config: cfg,
	}
	mux := api.mount()
	log.Fatal(api.run(mux))
}
