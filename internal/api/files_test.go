package api

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestGetFiles(t *testing.T) {
	// create a temporary directory to use as the prefix
	prefix := t.TempDir()

	// create a new file in the prefix
	err := os.WriteFile(path.Join(prefix, "test.txt"), []byte("test"), 0o644)
	assert.NoError(t, err)

	// create a new directory in the prefix
	err = os.MkdirAll(path.Join(prefix, "dir"), 0o777)
	assert.NoError(t, err)

	reqBody := GetFilesJSONBody{
		Path: string("/"),
	}
	reqBodyJSON, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/files", bytes.NewReader(reqBodyJSON))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server := NewServer(prefix)

	handler := http.HandlerFunc(NewStrictHandler(server, nil).GetFiles)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var files []File

	err = json.NewDecoder(rr.Body).Decode(&files)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(files))

	assert.Equal(t, "dir", files[0].Name)
	assert.Equal(t, true, files[0].IsDir)

	assert.Equal(t, "test.txt", files[1].Name)
	assert.Equal(t, false, files[1].IsDir)
}

func TestPostFiles(t *testing.T) {
	prefix := t.TempDir()

	reqBody0, err := json.Marshal(PostFilesJSONRequestBody{
		IsDir: true,
		Name:  "test",
	})
	assert.NoError(t, err)

	req0, err := http.NewRequest("POST", "/files", bytes.NewReader(reqBody0))
	assert.NoError(t, err)

	reqBody1, err := json.Marshal(PostFilesJSONRequestBody{
		IsDir: false,
		Name:  "test.txt",
	})
	assert.NoError(t, err)

	req1, err := http.NewRequest("POST", "/files", bytes.NewReader(reqBody1))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server := NewServer(prefix)

	handler := http.HandlerFunc(NewStrictHandler(server, nil).PostFiles)
	handler.ServeHTTP(rr, req0)
	assert.Equal(t, http.StatusNoContent, rr.Code)

	handler.ServeHTTP(rr, req1)
	assert.Equal(t, http.StatusNoContent, rr.Code)

	files, err := os.ReadDir(prefix)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(files))

	assert.Equal(t, "test", files[0].Name())
	assert.Equal(t, true, files[0].IsDir())

	assert.Equal(t, "test.txt", files[1].Name())
	assert.Equal(t, false, files[1].IsDir())
}

func TestDeleteFiles(t *testing.T) {
	prefix := t.TempDir()

	// create a new file in the prefix
	err := os.WriteFile(path.Join(prefix, "test.txt"), []byte("test"), 0o644)
	assert.NoError(t, err)

	// create a new directory in the prefix
	err = os.MkdirAll(path.Join(prefix, "test"), 0o777)
	assert.NoError(t, err)

	files, err := os.ReadDir(prefix)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(files))

	reqBody0, err := json.Marshal(DeleteFilesJSONRequestBody{
		Path: "test",
	})
	assert.NoError(t, err)

	req0, err := http.NewRequest("DELETE", "/files", bytes.NewReader(reqBody0))
	assert.NoError(t, err)

	reqBody1, err := json.Marshal(DeleteFilesJSONRequestBody{
		Path: "test.txt",
	})
	assert.NoError(t, err)

	req1, err := http.NewRequest("DELETE", "/files", bytes.NewReader(reqBody1))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server := NewServer(prefix)

	handler := http.HandlerFunc(NewStrictHandler(server, nil).DeleteFiles)
	handler.ServeHTTP(rr, req0)
	assert.Equal(t, http.StatusNoContent, rr.Code)

	handler.ServeHTTP(rr, req1)
	assert.Equal(t, http.StatusNoContent, rr.Code)

	files, err = os.ReadDir(prefix)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(files))

	// delete non-existing path
	reqBody2, err := json.Marshal(DeleteFilesJSONRequestBody{
		Path: "nonExisting.txt",
	})
	assert.NoError(t, err)
	req2, err := http.NewRequest("DELETE", "/files", bytes.NewReader(reqBody2))
	assert.NoError(t, err)
	rr2 := httptest.NewRecorder()
	handler.ServeHTTP(rr2, req2)
	assert.Equal(t, http.StatusNotFound, rr2.Code)
}
