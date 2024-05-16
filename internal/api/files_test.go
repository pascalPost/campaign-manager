package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileTree(t *testing.T) {
	tests := []struct {
		description        string
		url                string
		expectedStatusCode int
		checks             func(t *testing.T, rr *httptest.ResponseRecorder)
	}{
		{
			description:        "check root content",
			url:                "/fileTree",
			expectedStatusCode: http.StatusOK,
			checks: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var fileTreeEntry []FileTreeEntry

				err := json.NewDecoder(rr.Body).Decode(&fileTreeEntry)
				assert.NoError(t, err)

				assert.Equal(t, 2, len(fileTreeEntry))

				assert.Equal(t, "dir", fileTreeEntry[0].Name)
				assert.Equal(t, true, fileTreeEntry[0].IsDir)

				assert.Equal(t, "test.txt", fileTreeEntry[1].Name)
				assert.Equal(t, false, fileTreeEntry[1].IsDir)
			},
		},
		{
			description:        "check illegal root",
			url:                "/fileTree/",
			expectedStatusCode: http.StatusNotFound,
			checks:             func(t *testing.T, rr *httptest.ResponseRecorder) {},
		},
		{
			description:        "check nested folder content",
			url:                "/fileTree/dir",
			expectedStatusCode: http.StatusOK,
			checks: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var fileTreeEntry []FileTreeEntry

				err := json.NewDecoder(rr.Body).Decode(&fileTreeEntry)
				assert.NoError(t, err)

				assert.Equal(t, 1, len(fileTreeEntry))

				assert.Equal(t, "test_nested.txt", fileTreeEntry[0].Name)
				assert.Equal(t, false, fileTreeEntry[0].IsDir)
			},
		},
	}

	// create a temporary directory to use as the prefix
	prefix := t.TempDir()

	err := os.WriteFile(path.Join(prefix, "test.txt"), []byte("test"), 0o644)
	assert.NoError(t, err)

	err = os.MkdirAll(path.Join(prefix, "dir"), 0o777)
	assert.NoError(t, err)

	err = os.WriteFile(path.Join(prefix, "dir/test_nested.txt"), []byte("test2"), 0o644)
	assert.NoError(t, err)

	server := NewServer(prefix)
	handler := Handler(NewStrictHandler(server, nil))

	for _, test := range tests {
		t.Log(test.description)

		req, err := http.NewRequest("GET", test.url, nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, test.expectedStatusCode, rr.Code)

		test.checks(t, rr)
	}
}

func TestPostFileTree(t *testing.T) {
	tests := []struct {
		requestBody        PostFileTreeJSONRequestBody
		expectedStatusCode int
	}{
		{
			requestBody: PostFileTreeJSONRequestBody{
				IsDir: true,
				Name:  "test",
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			requestBody: PostFileTreeJSONRequestBody{
				IsDir: false,
				Name:  "test.txt",
			},
			expectedStatusCode: http.StatusNoContent,
		},
	}

	prefix := t.TempDir()
	server := NewServer(prefix)
	handler := Handler(NewStrictHandler(server, nil))

	for _, test := range tests {
		reqBody, err := json.Marshal(test.requestBody)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/fileTree", bytes.NewReader(reqBody))
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

//func TestDeleteFiles(t *testing.T) {
//	tests := []struct {
//		description        string
//		prepFunc           func(prefix string, t *testing.T)
//		requestBody        DeleteFilesJSONRequestBody
//		expectedStatusCode int
//		postTestCheck      func(prefix string, t *testing.T)
//	}{
//		{
//			description: "Test (empty) folder deletion",
//			prepFunc: func(prefix string, t *testing.T) {
//				// create a directory in the prefix
//				err := os.MkdirAll(path.Join(prefix, "test"), 0o777)
//				assert.NoError(t, err)
//
//				// check if folder was created
//				files, err := os.ReadDir(prefix)
//				assert.NoError(t, err)
//				assert.Equal(t, 1, len(files))
//				assert.Equal(t, "test", files[0].Name())
//				assert.Equal(t, true, files[0].IsDir())
//			},
//			requestBody: DeleteFilesJSONRequestBody{
//				Path: "test",
//			},
//			expectedStatusCode: http.StatusNoContent,
//			postTestCheck: func(prefix string, t *testing.T) {
//				files, err := os.ReadDir(prefix)
//				assert.NoError(t, err)
//				assert.Equal(t, 0, len(files), "The prefix folder should now be empty as the folder test was deleted.")
//			},
//		},
//		{
//			description: "Test file deletion",
//			prepFunc: func(prefix string, t *testing.T) {
//				// create a new file in the prefix
//				err := os.WriteFile(path.Join(prefix, "test.txt"), []byte("test"), 0o644)
//				assert.NoError(t, err)
//
//				// check if file was created
//				files, err := os.ReadDir(prefix)
//				assert.NoError(t, err)
//				assert.Equal(t, 1, len(files))
//				assert.Equal(t, "test.txt", files[0].Name())
//				assert.Equal(t, false, files[0].IsDir())
//			},
//			requestBody: DeleteFilesJSONRequestBody{
//				Path: "test.txt",
//			},
//			expectedStatusCode: http.StatusNoContent,
//			postTestCheck: func(prefix string, t *testing.T) {
//				files, err := os.ReadDir(prefix)
//				assert.NoError(t, err)
//				assert.Equal(t, 0, len(files), "The prefix folder should now be empty as the file test.txt was deleted.")
//			},
//		},
//		{
//			description: "Test non-existing path deletion",
//			prepFunc:    func(prefix string, t *testing.T) {},
//			requestBody: DeleteFilesJSONRequestBody{
//				Path: "nonExisting.txt",
//			},
//			expectedStatusCode: http.StatusNotFound,
//			postTestCheck:      func(prefix string, t *testing.T) {},
//		},
//	}
//
//	prefix := t.TempDir()
//
//	server := NewServer(prefix)
//	handler := Handler(NewStrictHandler(server, nil))
//
//	for _, test := range tests {
//		test.prepFunc(prefix, t)
//
//		reqBody, err := json.Marshal(test.requestBody)
//		assert.NoError(t, err)
//
//		req, err := http.NewRequest("DELETE", "/files", bytes.NewReader(reqBody))
//		assert.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		handler.ServeHTTP(rr, req)
//		assert.Equal(t, test.expectedStatusCode, rr.Code)
//
//		test.postTestCheck(prefix, t)
//	}
//}

func TestGetFileFilePath(t *testing.T) {
	tests := []struct {
		description string
		url         string
		checks      func(t *testing.T, rr *httptest.ResponseRecorder)
	}{
		{
			description: "test file download in root",
			url:         "/file/%2Ftest.txt",
			checks: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, "plain/text", rr.Header().Get("content-type"))

				body, err := io.ReadAll(rr.Body)
				assert.NoError(t, err)
				assert.Equal(t, "test", string(body))
			},
		},
		{
			description: "test file download in folder",
			url:         "/file/test%2Ftest_nested.txt",
			checks: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rr.Code)
			},
		},
	}

	prefix := t.TempDir()

	err := os.WriteFile(path.Join(prefix, "test.txt"), []byte("test"), 0o644)
	assert.NoError(t, err)

	err = os.MkdirAll(path.Join(prefix, "test"), 0o777)
	assert.NoError(t, err)

	err = os.WriteFile(path.Join(prefix, "test/test_nested.txt"), []byte("test2"), 0o644)
	assert.NoError(t, err)

	server := NewServer(prefix)
	handler := Handler(NewStrictHandler(server, nil))

	for _, test := range tests {
		t.Log(test.description)

		req, err := http.NewRequest("GET", test.url, nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		test.checks(t, rr)
	}

}

func TestPutFile(t *testing.T) {
	// create a temporary directory to use as the prefix
	prefix := t.TempDir()

	// create a new file in the prefix
	fileName := path.Join(prefix, "test.txt")
	err := os.WriteFile(fileName, []byte(""), 0o644)
	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", "/file/test.txt", strings.NewReader("test"))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	server := NewServer(prefix)

	handler := Handler(NewStrictHandler(server, nil))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	c, err := os.ReadFile(fileName)
	assert.NoError(t, err)
	assert.Equal(t, "test", string(c))
}

// TODO update testing structure to represent https://github.com/deepmap/oapi-codegen/blob/master/examples/petstore-expanded/strict/petstore_test.go
