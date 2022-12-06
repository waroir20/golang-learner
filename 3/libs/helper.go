package libs

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//---------------------------------------------------------------------------------------

// ExecuteTestRequest - Executes a fake API request (with options)
func ExecuteTestRequest(t *testing.T, router *gin.Engine, verb string, url string, code int, opts ...Option) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	var bodyStr string
	cfg := &config{}
	for _, o := range opts {
		o(cfg)
	}
	if cfg.body != nil && cfg.body != struct{}{} {
		bdy, err := json.Marshal(cfg.body)
		if err != nil {
			t.Errorf("Error marshalling object: %s", err)
		} else {
			bodyStr = string(bdy)
		}
	}
	req, err := http.NewRequest(verb, url, strings.NewReader(bodyStr))
	if err != nil {
		t.Errorf("Error creating new request: %s", err)
	}
	if cfg.headers != nil {
		header := http.Header{}
		for k, v := range cfg.headers {
			header.Add(k, v)
		}
		req.Header = header
	}
	router.ServeHTTP(rr, req)
	checkStatus(t, rr, code)
	return rr
}

func checkStatus(t *testing.T, rr *httptest.ResponseRecorder, code int) {
	if status := rr.Code; status != code {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, code)
	}
}
