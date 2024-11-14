package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetDateTimeHandler(t *testing.T) {
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

		request := httptest.NewRequest(http.MethodGet, "/datetime", nil)
		responseRecorder := httptest.NewRecorder()

		GetDateTimeHandler(responseRecorder, request)

		if responseRecorder.Code != http.StatusOK {
			t.Errorf("Want status '%d', got '%d'", http.StatusOK, responseRecorder.Code)
		}

		err := json.NewDecoder(responseRecorder.Body).Decode(&returnedDateTime)

		assert.NoError(t, err)

		if !slices.Contains(possibleValues, returnedDateTime.Message) {
			t.Errorf("Want '%s' or '%s' or '%s', got '%s'", possibleValues[0], possibleValues[1], possibleValues[2], returnedDateTime)
		}
	})
}
