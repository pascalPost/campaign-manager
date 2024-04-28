package api

import (
	"context"
	"os"
	"path"
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
