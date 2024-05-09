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

func NewFilesService(prefix string) *FilesService {
	return &FilesService{
		Prefix: prefix,
	}
}

func (s *FilesService) GetFiles(ctx context.Context, request GetFilesRequestObject) (GetFilesResponseObject, error) {
	p := path.Join(s.Prefix, request.Body.Path)

	files, err := os.ReadDir(p)
	if err != nil {
		return nil, err
	}

	result := make([]File, 0, len(files))
	for _, f := range files {
		result = append(result, File{
			Name:  f.Name(),
			IsDir: f.IsDir(),
		})
	}

	return GetFiles200JSONResponse(result), nil
}

func (s *FilesService) PostFiles(ctx context.Context, request PostFilesRequestObject) (PostFilesResponseObject, error) {
	name := path.Join(s.Prefix, request.Body.Name)
	isDir := request.Body.IsDir

	_, err := os.Stat(name)
	if err != nil && !os.IsNotExist(err) {
		// already exists
		return nil, err
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

	return PostFiles204Response{}, nil
}

func (s *FilesService) DeleteFiles(ctx context.Context, request DeleteFilesRequestObject) (DeleteFilesResponseObject, error) {
	p := path.Join(s.Prefix, request.Body.Path)

	_, err := os.Stat(p)
	if err != nil {
		if os.IsNotExist(err) {
			return DeleteFiles404Response{}, nil
		}

		return nil, err
	}

	err = os.RemoveAll(p)
	if err != nil {
		return nil, err
	}

	return DeleteFiles204Response{}, nil
}

func (s *FilesService) PutFiles(ctx context.Context, request PutFilesRequestObject) (PutFilesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *FilesService) GetFileFilePath(ctx context.Context, request GetFileFilePathRequestObject) (GetFileFilePathResponseObject, error) {
	filePath := filepath.Join(s.Prefix, request.FilePath)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Error("GetFile", "filePath", filePath, "Path not exists response", http.StatusNotFound)
			return GetFileFilePath404Response{}, nil
		}

		slog.Error("GetFile", "filePath", filePath, "err", err)
		return nil, err
	}

	if fileInfo.IsDir() {
		slog.Error("GetFile", "filePath", filePath, "isDir response", http.StatusBadRequest)
		return GetFileFilePath400Response{}, nil
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileType := http.DetectContentType(content)
	if !strings.HasPrefix(fileType, "text/plain") {
		slog.Error("GetFile", "filePath", filePath, "file type", fileType, "non plain/text file response", http.StatusBadRequest)
		return GetFileFilePath400Response{}, nil
	}

	slog.Debug("GetFile", "filePath", filePath, "response", http.StatusOK)
	return GetFileFilePath200PlaintextResponse{
		Body: bytes.NewReader(content),
	}, nil
}

func (s *FilesService) PutFileFilePath(ctx context.Context, request PutFileFilePathRequestObject) (PutFileFilePathResponseObject, error) {
	filePath := filepath.Join(s.Prefix, request.FilePath)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Error("GetFile", "filePath", filePath, "Path not exists response", http.StatusNotFound)
			return PutFileFilePath404Response{}, nil
		}

		slog.Error("PutFile", "filePath", filePath, "err", err)
		return nil, err
	}

	if fileInfo.IsDir() {
		slog.Error("PutFile", "filePath", filePath, "isDir response", http.StatusBadRequest)
		return PutFileFilePath400Response{}, nil
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
