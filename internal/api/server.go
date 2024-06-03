//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../api/openapi.yaml

package api

import "context"

type Server struct {
}

var _ StrictServerInterface = (*Server)(nil)

func NewServer() *Server {
	return &Server{}
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
