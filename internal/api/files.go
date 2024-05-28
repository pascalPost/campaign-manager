package api

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FilesService struct {
	Prefix string
}

const NonLocalPathMessage = "Non local path."
const PathNotFoundMessage = "Path not found."
const NotPlainTextFileMessage = "Not plain/text file."

func nonLocalPathResponse() NonLocalPathJSONResponse {
	return NonLocalPathJSONResponse{Message: NonLocalPathMessage}
}

func pathNotFoundResponse() PathNotFoundJSONResponse {
	return PathNotFoundJSONResponse{Message: PathNotFoundMessage}
}

// isSavePath checks the given path to prevent path traversal attacks
func isSavePath(path string) bool {
	return filepath.IsLocal(path)
}

// normalizePath transforms the abs. or rel. path into a relative path
func normalizePath(path string) string {
	return filepath.Clean(filepath.Join("./", path))
}

func NewFilesService(prefix string) *FilesService {
	return &FilesService{
		Prefix: prefix,
	}
}

func getFileTree(prefix string, filePath string) ([]FileTreeEntry, error) {
	rootPath := filepath.Join(prefix, filePath)

	files, err := os.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}

	result := make([]FileTreeEntry, 0, len(files))
	for _, f := range files {
		result = append(result, FileTreeEntry{
			Path:  filepath.Join("/", filePath, f.Name()),
			IsDir: f.IsDir(),
		})
	}
	return result, nil
}

func (s *FilesService) GetFileTree(ctx context.Context, request GetFileTreeRequestObject) (GetFileTreeResponseObject, error) {
	result, err := getFileTree(s.Prefix, "/")
	if err != nil {
		return nil, err
	}

	return GetFileTree200JSONResponse(result), nil
}

func (s *FilesService) GetFileTreePath(ctx context.Context, request GetFileTreePathRequestObject) (GetFileTreePathResponseObject, error) {
	normPath := normalizePath(request.Path)
	if !isSavePath(normPath) {
		return GetFileTreePath400JSONResponse{nonLocalPathResponse()}, nil
	}

	result, err := getFileTree(s.Prefix, normPath)
	if err != nil {
		return nil, err
	}
	return GetFileTreePath200JSONResponse(result), nil
}

func (s *FilesService) PostFileTree(ctx context.Context, request PostFileTreeRequestObject) (PostFileTreeResponseObject, error) {
	reqPath := normalizePath(request.Body.Path)
	if !isSavePath(reqPath) {
		return PostFileTree400JSONResponse{nonLocalPathResponse()}, nil
	}

	name := path.Join(s.Prefix, reqPath)
	isDir := request.Body.IsDir

	_, err := os.Stat(name)
	if err == nil {
		return PostFileTree409JSONResponse{
			Message: "Path already exists.",
		}, nil
	}

	if isDir {
		err := os.MkdirAll(name, os.ModePerm)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := os.Create(name)
		if err != nil {
			return nil, err
		}
	}

	return PostFileTree201JSONResponse{
		Path: filepath.Join("/", reqPath),
	}, nil
}

func (s *FilesService) DeleteFileTreePath(ctx context.Context, request DeleteFileTreePathRequestObject) (DeleteFileTreePathResponseObject, error) {
	normPath := normalizePath(request.Path)
	if !isSavePath(normPath) {
		return DeleteFileTreePath400JSONResponse{nonLocalPathResponse()}, nil
	}

	reqPath := filepath.Join(s.Prefix, normPath)

	_, err := os.Stat(reqPath)
	if err != nil {
		if os.IsNotExist(err) {
			return DeleteFileTreePath404JSONResponse{pathNotFoundResponse()}, nil
		}

		return nil, err
	}

	err = os.RemoveAll(reqPath)
	if err != nil {
		return nil, err
	}

	return DeleteFileTreePath200JSONResponse{
		Path: filepath.Join("/", request.Path),
	}, nil
}

//func (s *FilesService) PutFiles(ctx context.Context, request PutFileTreeRequestObject) (PutFilesResponseObject, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (s *FilesService) GetFileFilePath(ctx context.Context, request GetFileFilePathRequestObject) (GetFileFilePathResponseObject, error) {
	normPath := normalizePath(request.FilePath)
	if !isSavePath(normPath) {
		return GetFileFilePath400JSONResponse{BadRequestJSONResponse{NonLocalPathMessage}}, nil
	}

	filePath := filepath.Join(s.Prefix, normPath)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return GetFileFilePath404JSONResponse{pathNotFoundResponse()}, nil
		}

		slog.Error("GetFile", "filePath", filePath, "err", err)
		return nil, err
	}

	if fileInfo.IsDir() {
		return GetFileFilePath400JSONResponse{BadRequestJSONResponse{NotPlainTextFileMessage}}, nil
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileType := http.DetectContentType(content)
	if !strings.HasPrefix(fileType, "text/plain") {
		slog.Error("GetFile", "filePath", filePath, "file type", fileType, "non plain/text file response", http.StatusBadRequest)
		return GetFileFilePath400JSONResponse{BadRequestJSONResponse{NotPlainTextFileMessage}}, nil
	}

	slog.Debug("GetFile", "filePath", filePath, "response", http.StatusOK)
	return GetFileFilePath200PlaintextResponse{
		Body: bytes.NewReader(content),
	}, nil
}

func (s *FilesService) PutFileFilePath(ctx context.Context, request PutFileFilePathRequestObject) (PutFileFilePathResponseObject, error) {
	normPath := normalizePath(request.FilePath)
	if !isSavePath(normPath) {
		return PutFileFilePath400JSONResponse{BadRequestJSONResponse{NonLocalPathMessage}}, nil
	}

	filePath := filepath.Join(s.Prefix, normPath)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return PutFileFilePath404JSONResponse{pathNotFoundResponse()}, nil
		}

		slog.Error("PutFile", "filePath", filePath, "err", err)
		return nil, err
	}

	if fileInfo.IsDir() {
		slog.Error("PutFile", "filePath", filePath, "isDir response", http.StatusBadRequest)
		return PutFileFilePath400JSONResponse{BadRequestJSONResponse{NotPlainTextFileMessage}}, nil
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		slog.Error("PutFile", "filePath", filePath, "ReadAll body err", err)
		return nil, err
	}

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0o644)
	if err != nil {
		slog.Error("PutFile", "filePath", filePath, "OpenFile err", err)
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			slog.Error("PutFile", "filePath", filePath, "Close file err", err)
		}
	}()

	_, err = f.Write(body)
	if err != nil {
		slog.Error("PutFile", "filePath", filePath, "Write file err", err)
		return nil, err
	}

	return PutFileFilePath200Response{}, nil
}
