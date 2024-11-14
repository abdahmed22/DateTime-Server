package ginserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type JSONResponse struct {
	Message string `json:"message"`
}

func TestPingRoute(t *testing.T) {
	now := time.Now()

	firstPossibleValue := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second()-1,
		now.Nanosecond(), now.Location()).Format("2006-01-02 15:04")

	secondPossibleValue := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second()-2,
		now.Nanosecond(), now.Location()).Format("2006-01-02 15:04")

	thirdPossibleValue := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second()-3,
		now.Nanosecond(), now.Location()).Format("2006-01-02 15:04")

	possibleValues := []string{firstPossibleValue, secondPossibleValue, thirdPossibleValue}

	var returnedDateTime JSONResponse
	t.Run("happy path - can hit the api and return date & time", func(*testing.T) {
		router := SetupRouter()

		responseRecorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/datetime", nil)
		router.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusOK {
			t.Errorf("Want status '%d', got '%d'", http.StatusOK, responseRecorder.Code)
		}

		err := json.NewDecoder(responseRecorder.Body).Decode(&returnedDateTime)

		assert.NoError(t, err)

		if !slices.Contains(possibleValues, returnedDateTime.Message) {
			t.Errorf("Want '%s' or '%s' or '%s', got '%s'", possibleValues[0], possibleValues[1], possibleValues[2], returnedDateTime)
		}
	})

	t.Run("sad path - unavailable end point", func(*testing.T) {
		router := SetupRouter()

		responseRecorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/date", nil)
		router.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusNotFound {
			t.Errorf("Want status '%d', got '%d'", http.StatusNotFound, responseRecorder.Code)
		}

		assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	})

	t.Run("sad path - unavailable method", func(*testing.T) {
		router := SetupRouter()

		responseRecorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/datetime", nil)
		router.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusNotFound {
			t.Errorf("Want status '%d', got '%d'", http.StatusNotFound, responseRecorder.Code)
		}

		assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	})
}
