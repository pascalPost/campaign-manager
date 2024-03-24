package routes

import (
	"github.com/campaign-manager/src"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	newRecorder := httptest.NewRecorder()
	cm.Routes().ServeHTTP(newRecorder, req)

	assert.Equal(t, 200, newRecorder.Code)
	assert.Equal(t, "OK", newRecorder.Body.String())
}
