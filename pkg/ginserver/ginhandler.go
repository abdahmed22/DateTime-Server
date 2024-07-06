package ginserver

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GetDateTimeHandler(c *gin.Context) {
	currentDateTime := time.Now()

	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{
		"message": currentDateTime.Format("2006-01-02 15:04"),
	})
}
