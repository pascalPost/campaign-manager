package project

import (
	"github.com/rs/xid"
	"os"
	"path/filepath"
)

type Storage interface {
	GetProjects() []*Project
	AddProject(*Project)
}

type Project struct {
	Id   string
	Name string
	Path string
}

func NewProject(name string, rootFolder string) (*Project, error) {
	id := xid.New().String()

	path := filepath.Join(rootFolder, id)

	if err := createProjectFolder(path); err != nil {
		return nil, err
	}

	return &Project{
		Id:   id,
		Name: name,
		Path: path,
	}, nil
}

func createProjectFolder(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
