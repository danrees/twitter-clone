package main

import (
	"log"
	"os"

	"github.com/danrees/twitter-clone/router"
)

func main() {

	addr, ok := os.LookupEnv("SERVER_ADDR")
	if !ok {
		addr = ":8080"
	}

	r := router.New()
	log.Printf("Starting server on %v", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
