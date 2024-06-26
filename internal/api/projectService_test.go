package api

import (
	"bytes"
	"encoding/json"
	"github.com/campaign-manager/internal/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestPostProjects(t *testing.T) {
	prefix := t.TempDir()

	reqBody, err := json.Marshal(PostProjectsJSONRequestBody{Name: "test"})
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/projects", bytes.NewReader(reqBody))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	store := storage.NewInMemStorage()
	server := NewServer(prefix, store)

	handler := Handler(NewStrictHandler(server, nil))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	projects := store.GetProjects()
	assert.Equal(t, 1, len(projects))
	assert.Equal(t, "test", projects[0].Name)

	files, err := os.ReadDir(prefix)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(files))
	assert.Equal(t, projects[0].Id, files[0].Name())
}
