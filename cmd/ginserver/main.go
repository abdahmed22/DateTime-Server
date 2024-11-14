package main

import (
	"log"

	"github.com/codescalersinternships/DateTime-Server-Abdelrahman-Mahmoud/pkg/ginserver"
)

func main() {
	r := ginserver.SetupRouter()

	err := r.Run(":8080")

	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
