package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type JSONResponse struct {
	Message string `json:"message"`
}

func GetDateTimeHandler(w http.ResponseWriter, rc *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	currentDateTime := time.Now().Format("2006-01-02 15:04")

	responseMessage := JSONResponse{currentDateTime}

	err := json.NewEncoder(w).Encode(&responseMessage)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
