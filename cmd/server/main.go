package main

import (
	"log"

	"github.com/rkiii-code/discord-demo/internal/server"
)

func main() {
	r := server.New() // gin.Engine を返す
	log.Println("listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
