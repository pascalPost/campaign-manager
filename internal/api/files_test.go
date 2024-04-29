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
	tests := []struct {
		requestBody        PostFilesJSONRequestBody
		expectedStatusCode int
	}{
		{
			requestBody: PostFilesJSONRequestBody{
				IsDir: true,
				Name:  "test",
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			requestBody: PostFilesJSONRequestBody{
				IsDir: false,
				Name:  "test.txt",
			},
			expectedStatusCode: http.StatusNoContent,
		},
	}

	prefix := t.TempDir()
	server := NewServer(prefix)
	handler := http.HandlerFunc(NewStrictHandler(server, nil).PostFiles)

	for _, test := range tests {
		reqBody, err := json.Marshal(test.requestBody)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/files", bytes.NewReader(reqBody))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, test.expectedStatusCode, rr.Code)
	}

	files, err := os.ReadDir(prefix)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(files))

	assert.Equal(t, "test", files[0].Name())
	assert.Equal(t, true, files[0].IsDir())

	assert.Equal(t, "test.txt", files[1].Name())
	assert.Equal(t, false, files[1].IsDir())
}

func TestDeleteFiles(t *testing.T) {
	tests := []struct {
		description        string
		prepFunc           func(prefix string, t *testing.T)
		requestBody        DeleteFilesJSONRequestBody
		expectedStatusCode int
		postTestCheck      func(prefix string, t *testing.T)
	}{
		{
			description: "Test (empty) folder deletion",
			prepFunc: func(prefix string, t *testing.T) {
				// create a directory in the prefix
				err := os.MkdirAll(path.Join(prefix, "test"), 0o777)
				assert.NoError(t, err)

				// check if folder was created
				files, err := os.ReadDir(prefix)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(files))
				assert.Equal(t, "test", files[0].Name())
				assert.Equal(t, true, files[0].IsDir())
			},
			requestBody: DeleteFilesJSONRequestBody{
				Path: "test",
			},
			expectedStatusCode: http.StatusNoContent,
			postTestCheck: func(prefix string, t *testing.T) {
				files, err := os.ReadDir(prefix)
				assert.NoError(t, err)
				assert.Equal(t, 0, len(files), "The prefix folder should now be empty as the folder test was deleted.")
			},
		},
		{
			description: "Test file deletion",
			prepFunc: func(prefix string, t *testing.T) {
				// create a new file in the prefix
				err := os.WriteFile(path.Join(prefix, "test.txt"), []byte("test"), 0o644)
				assert.NoError(t, err)

				// check if file was created
				files, err := os.ReadDir(prefix)
				assert.NoError(t, err)
				assert.Equal(t, 1, len(files))
				assert.Equal(t, "test.txt", files[0].Name())
				assert.Equal(t, false, files[0].IsDir())
			},
			requestBody: DeleteFilesJSONRequestBody{
				Path: "test.txt",
			},
			expectedStatusCode: http.StatusNoContent,
			postTestCheck: func(prefix string, t *testing.T) {
				files, err := os.ReadDir(prefix)
				assert.NoError(t, err)
				assert.Equal(t, 0, len(files), "The prefix folder should now be empty as the file test.txt was deleted.")
			},
		},
		{
			description: "Test non-existing path deletion",
			prepFunc:    func(prefix string, t *testing.T) {},
			requestBody: DeleteFilesJSONRequestBody{
				Path: "nonExisting.txt",
			},
			expectedStatusCode: http.StatusNotFound,
			postTestCheck:      func(prefix string, t *testing.T) {},
		},
	}

	prefix := t.TempDir()

	server := NewServer(prefix)
	handler := http.HandlerFunc(NewStrictHandler(server, nil).DeleteFiles)

	for _, test := range tests {
		test.prepFunc(prefix, t)

		reqBody, err := json.Marshal(test.requestBody)
		assert.NoError(t, err)

		req, err := http.NewRequest("DELETE", "/files", bytes.NewReader(reqBody))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.expectedStatusCode, rr.Code)

		test.postTestCheck(prefix, t)
	}
}
