package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	api_key := os.Getenv("API_KEY")
	cfg := config{
		addr: ":8080",
	}

	api := &api{
		config: cfg,
	}
	mux := api.mount()
	log.Fatal(api.run(mux))
}
