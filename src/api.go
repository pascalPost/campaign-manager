package cm

import (
	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"context"
	"database/sql"
	cmdb "github.com/campaign-manager/src/db"
	_ "github.com/campaign-manager/src/proto/cm/v1"
	protocmv1 "github.com/campaign-manager/src/proto/cm/v1"
	"github.com/campaign-manager/src/proto/cm/v1/protocmv1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type CampaignManagerService struct {
	protocmv1connect.UnimplementedCampaignManagerServiceHandler
	db *sql.DB
}

func NewCampaignManagerService(db *sql.DB) *CampaignManagerService {
	return &CampaignManagerService{db: db}
}

func (s *CampaignManagerService) Ping(
	ctx context.Context,
	req *connect.Request[protocmv1.PingRequest],
) (*connect.Response[protocmv1.PingResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&protocmv1.PingResponse{
		Message: "OK",
	})
	return res, nil
}

func (s *CampaignManagerService) NewProject(
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

func (s *CampaignManagerService) GetSettings(
	ctx context.Context,
	req *connect.Request[protocmv1.GetSettingsRequest],
) (*connect.Response[protocmv1.GetSettingsResponse], error) {

	settings, err := cmdb.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if settings == nil {
		res := connect.NewResponse(&protocmv1.GetSettingsResponse{
			WorkingDir:  "",
			LsfUsername: "",
			LsfPassword: "",
		})

		return res, nil
	}

	res := connect.NewResponse(&protocmv1.GetSettingsResponse{
		WorkingDir:  settings.WorkingDir(),
		LsfUsername: settings.LSFUsername(),
		LsfPassword: settings.LSFPassword(),
	})
	return res, nil
}

func (s *CampaignManagerService) SetSettings(
	ctx context.Context,
	req *connect.Request[protocmv1.SetSettingsRequest],
) (*connect.Response[protocmv1.SetSettingsResponse], error) {
	log.Println("Request headers: ", req.Header())
	log.Println("Request Msg: ", req.Msg)

	workingDir := req.Msg.WorkingDir

	// TODO save workingDir to db

	res := connect.NewResponse(&protocmv1.SetSettingsResponse{
		WorkingDir: workingDir,
	})
	return res, nil
}

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}

func Server() {
	db := cmdb.ConnectDB()
	defer cmdb.DisconnectDB(db)

	mux := http.NewServeMux()
	mux.Handle(protocmv1connect.NewCampaignManagerServiceHandler(NewCampaignManagerService(db)))
	err := http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		withCORS(h2c.NewHandler(mux, &http2.Server{})),
	)
	if err != nil {
		log.Fatal(err)
	}
}
