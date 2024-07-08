package main

import (
	"log"

	"github.com/codescalersinternships/DateTime-Server-Abdelrahman-Mahmoud/pkg/ginserver"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/datetime", ginserver.GetDateTimeHandler)

	err := r.Run()

	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
