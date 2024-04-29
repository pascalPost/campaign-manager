//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../api/openapi.yaml

package api

import "context"

type Server struct {
	*FilesService
}

var _ StrictServerInterface = (*Server)(nil)

func NewServer(prefix string) *Server {
	filesService := NewFilesService(prefix)
	return &Server{
		FilesService: filesService,
	}
}

// Ping server
// (GET /ping)
func (s *Server) GetPing(ctx context.Context, request GetPingRequestObject) (GetPingResponseObject, error) {
	return nil, nil
}

// List projects
// (GET /projects)
func (s *Server) GetProjects(ctx context.Context, request GetProjectsRequestObject) (GetProjectsResponseObject, error) {
	return nil, nil
}

// Add new projects
// (POST /projects)
func (s *Server) PostProjects(ctx context.Context, request PostProjectsRequestObject) (PostProjectsResponseObject, error) {
	return nil, nil
}

// List tasks
// (GET /tasks)
func (s *Server) GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error) {
	return nil, nil
}

// Add new task
// (POST /tasks)
func (s *Server) PostTasks(ctx context.Context, request PostTasksRequestObject) (PostTasksResponseObject, error) {
	return nil, nil
}
