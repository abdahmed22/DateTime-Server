package main

import (
	"github.com/codescalersinternships/DateTime-Server-Abdelrahman-Mahmoud/pkg/ginserver"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/datetime", ginserver.GetDateTimeHandler)

	r.Run()
}
