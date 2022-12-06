package api_test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	api "golang-learner/3"
	"golang-learner/3/libs"
	"net/http"
	"os"
	"testing"
)

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//---------------------------------------------------------------------------------------

const baseURL = "/api/v1/samples"

var (
	router *gin.Engine
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	// Create Routes
	router = gin.Default()
	api.RegisterRoutes(router)
}
func cleanup() {
	log.Debug().Msg("In cleanup")
	// TODO any post test cleanup
}
func shutdown() {}

func TestGet(t *testing.T) {
	t.Cleanup(cleanup) // This is run at the end of the test and all of its subtests
	t.Run("API returns 200 and expected phrase", func(t *testing.T) {
		resp := libs.ExecuteTestRequest(t, router, "GET", baseURL, http.StatusOK)
		assert.Equal(t, "Where no Lopporit has gone before!", resp.Body.String())
	})
}

func TestPost(t *testing.T) {
	t.Cleanup(cleanup)
	t.Run("API returns 201 and a copy of original request data", func(t *testing.T) {
		requestPayload := api.Aggregate{
			Id:   "61e3ef7d-6af9-4105-a84f-d671b916ab60",
			Name: "Sleepingway",
			Data: "Livingway's is the best ways",
		}
		resp := libs.ExecuteTestRequest(t, router, "POST", baseURL, http.StatusCreated, libs.WithBody(requestPayload))
		var responseBody api.Aggregate
		err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
		assert.Nil(t, err, "error during unmarshal")
		assert.Equal(t, requestPayload, responseBody)
	})
	t.Run("API returns 400 when request payload fails validation", func(t *testing.T) {
		requestPayload := api.Aggregate{
			Id:   "I am not a UUID",
			Name: "Runningway",
			Data: "Bestway's Burrows",
		}
		_ = libs.ExecuteTestRequest(t, router, "POST", baseURL, http.StatusBadRequest, libs.WithBody(requestPayload))
	})
}

func TestPut(t *testing.T) {
	t.Cleanup(cleanup)
	headers := make(map[string]string)
	headers["data"] = "Livingway's is the best ways"
	headers["name"] = "Sleepingway"
	t.Run("API returns 202 and response body built from path and headers", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", baseURL, "61e3ef7d-6af9-4105-a84f-d671b916ab60")
		resp := libs.ExecuteTestRequest(t, router, "PUT", url, http.StatusAccepted, libs.WithHeaders(headers))
		var responseBody api.Aggregate
		err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
		assert.Nil(t, err, "error during unmarshal")
		expectedPayload := api.Aggregate{
			Id:   "61e3ef7d-6af9-4105-a84f-d671b916ab60",
			Name: "Sleepingway",
			Data: "Livingway's is the best ways",
		}
		assert.Equal(t, expectedPayload, responseBody)
	})
	t.Run("API returns 500 when path parameter empty", func(t *testing.T) { //The goal here is to test out different response statuses
		url := fmt.Sprintf("%s/", baseURL)
		_ = libs.ExecuteTestRequest(t, router, "PUT", url, http.StatusInternalServerError, libs.WithHeaders(headers))
	})
}

func TestDelete(t *testing.T) {
	t.Cleanup(cleanup)
	t.Run("API returns 418 and a copy of original request data", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", baseURL, "61e3ef7d-6af9-4105-a84f-d671b916ab60")
		_ = libs.ExecuteTestRequest(t, router, "DELETE", url, http.StatusTeapot)
	})
	t.Run("API returns 422 when path parameter is 'Livingway'", func(t *testing.T) { //The goal here is to test out different response statuses
		url := fmt.Sprintf("%s/%s", baseURL, "Livingway")
		_ = libs.ExecuteTestRequest(t, router, "DELETE", url, http.StatusUnprocessableEntity)
	})
}
