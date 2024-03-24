package cm

import (
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
	Routes().ServeHTTP(newRecorder, req)

	assert.Equal(t, 200, newRecorder.Code)
	assert.Equal(t, "OK", newRecorder.Body.String())
}
