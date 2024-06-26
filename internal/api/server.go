//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../api/openapi.yaml

package api

import (
	"context"
	"github.com/campaign-manager/internal/project"
	"github.com/campaign-manager/internal/storage"
	"strings"
)

type Server struct {
	storage      storage.Storage
	filePathRoot string
}

var _ StrictServerInterface = (*Server)(nil)

func NewServer(filePathRoot string, s storage.Storage) *Server {
	return &Server{storage: s, filePathRoot: filePathRoot}
}

// Ping server
// (GET /ping)
func (s *Server) GetPing(ctx context.Context, request GetPingRequestObject) (GetPingResponseObject, error) {
	return nil, nil
}

func removeFilePathRoot(path string, root string) string {
	// remove prefix for usage with the file system service
	return strings.TrimPrefix(path, root)
}

// List projects
// (GET /projects)
func (s *Server) GetProjects(ctx context.Context, request GetProjectsRequestObject) (GetProjectsResponseObject, error) {
	store := s.storage.GetProjects()

	projects := make([]Project, len(store))

	for _, p := range store {
		path := removeFilePathRoot(p.Path, s.filePathRoot)

		projects = append(projects, Project{
			Id:   p.Id,
			Name: p.Name,
			Path: path,
		})
	}

	return GetProjects200JSONResponse(projects), nil
}

// Add new project
// (POST /projects)
func (s *Server) PostProjects(ctx context.Context, request PostProjectsRequestObject) (PostProjectsResponseObject, error) {
	name := request.Body.Name
	p, err := project.NewProject(name, s.filePathRoot)
	if err != nil {
		return nil, err
	}

	s.storage.AddProject(p)

	return PostProjects201JSONResponse{Id: p.Id}, nil
}

// List jobs
// (GET /jobs)
func (s *Server) GetJobs(ctx context.Context, request GetJobsRequestObject) (GetJobsResponseObject, error) {
	return nil, nil
}

// Add new job
// (POST /jobs)
func (s *Server) PostJobs(ctx context.Context, request PostJobsRequestObject) (PostJobsResponseObject, error) {
	return nil, nil
}
