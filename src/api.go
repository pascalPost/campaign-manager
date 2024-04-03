package cm

import (
	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"context"
	"database/sql"
	_ "github.com/campaign-manager/src/proto/cm/v1"
	protocmv1 "github.com/campaign-manager/src/proto/cm/v1"
	"github.com/campaign-manager/src/proto/cm/v1/protocmv1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type PingServer struct {
	protocmv1connect.UnimplementedPingServiceHandler
}

func (s *PingServer) Ping(
	ctx context.Context,
	req *connect.Request[protocmv1.PingRequest],
) (*connect.Response[protocmv1.PingResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&protocmv1.PingResponse{
		Message: "OK",
	})
	//res.Header().Set("Greet-Version", "v1")
	return res, nil
}

type ProjectServer struct {
	protocmv1connect.UnimplementedNewProjectServiceHandler
	db *sql.DB
}

func NewProjectServer(db *sql.DB) *ProjectServer {
	return &ProjectServer{db: db}
}

func (s *ProjectServer) NewProject(
	ctx context.Context,
	req *connect.Request[protocmv1.NewProjectRequest],
) (*connect.Response[protocmv1.NewProjectResponse], error) {
	log.Println("Request headers: ", req.Header())
	log.Println("Request Msg: ", req.Msg)

	projectId := uint64(11)

	res := connect.NewResponse(&protocmv1.NewProjectResponse{
		ProjectId: &projectId,
	})
	return res, nil
}

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}

func Server() {
	db := ConnectDB()

	mux := http.NewServeMux()
	mux.Handle(protocmv1connect.NewPingServiceHandler(&PingServer{}))
	mux.Handle(protocmv1connect.NewNewProjectServiceHandler(NewProjectServer(db)))
	err := http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		withCORS(h2c.NewHandler(mux, &http2.Server{})),
	)
	if err != nil {
		log.Fatal(err)
	}
}
