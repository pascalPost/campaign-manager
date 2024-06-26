package storage

import "github.com/campaign-manager/internal/project"

type Storage interface {
	project.Storage
}

type InMemStorage struct {
	projects []*project.Project
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{}
}

func (s *InMemStorage) GetProjects() []*project.Project {
	return s.projects
}

func (s *InMemStorage) AddProject(project *project.Project) {
	s.projects = append(s.projects, project)
}
