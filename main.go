package main

import (
	"log"
	"os"

	"github.com/danrees/twitter-clone/router"
	"github.com/gin-gonic/gin"
)

func setup() (string, *gin.Engine) {
	addr, ok := os.LookupEnv("SERVER_ADDR")
	if !ok {
		addr = ":8080"
	}

	r := router.New()
	return addr, r
}

func main() {

	addr, r := setup()
	log.Printf("Starting server on %v", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
