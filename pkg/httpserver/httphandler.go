package httpserver

import (
	"fmt"
	"net/http"
	"time"
)

type JSONResponse struct {
	Message string `json:"message"`
}

func GetDateTimeHandler(w http.ResponseWriter, rc *http.Request) {
	currentDateTime := time.Now().Format("2006-01-02 15:04")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, currentDateTime)

}
