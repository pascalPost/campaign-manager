package cm

import (
	"connectrpc.com/connect"
	"context"
	"database/sql"
	cmdb "github.com/campaign-manager/src/db"
	"github.com/campaign-manager/src/lsf"
	_ "github.com/campaign-manager/src/proto/cm/v1"
	protocmv1 "github.com/campaign-manager/src/proto/cm/v1"
	"github.com/campaign-manager/src/proto/cm/v1/protocmv1connect"
	"github.com/campaign-manager/src/types"
	"github.com/rs/cors"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
	"github.com/swaggest/usecase"
	"log"
	"net/http"
	"os"
)

type CampaignManagerService struct {
	protocmv1connect.UnimplementedCampaignManagerServiceHandler
	db *sql.DB
}

func NewCampaignManagerService(db *sql.DB) *CampaignManagerService {
	return &CampaignManagerService{db: db}
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
	_ context.Context,
	req *connect.Request[protocmv1.SetSettingsRequest],
) (*connect.Response[protocmv1.SetSettingsResponse], error) {

	settings := types.NewSettings(req.Msg.WorkingDir, req.Msg.LsfUsername, req.Msg.LsfPassword)
	err := cmdb.SetSettings(s.db, settings)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&protocmv1.SetSettingsResponse{})
	return res, nil
}

func (s *CampaignManagerService) GetLsfJobs(
	_ context.Context,
	_ *connect.Request[protocmv1.GetLsfJobsRequest],
) (*connect.Response[protocmv1.GetLsfJobsResponse], error) {
	jobs := lsf.Jobs()

	jobs_res := make([]*protocmv1.Job, 0)

	for _, job := range jobs {
		jobs_res = append(jobs_res, &protocmv1.Job{
			Command:   job.Command,
			ExHosts:   job.ExHosts,
			FromHost:  job.FromHost,
			JobId:     job.JobId,
			JobName:   job.JobName,
			JobStatus: job.JobStatus,
			Queue:     job.Queue,
			//SubmitTime: job.SubmitTime,
			User: job.User,
		})
	}

	res := connect.NewResponse(&protocmv1.GetLsfJobsResponse{
		Jobs: jobs_res,
	})

	return res, nil
}

type pingInput struct{}

type pingOutput struct {
	Message string `json:"message"`
}

type lsfJobInput struct{}

type lsfJobOutput struct {
	Jobs []lsf.Job `json:"jobs"`
}

func Server() {
	db := cmdb.ConnectDB()
	defer cmdb.DisconnectDB(db)

	// openapi
	r := openapi31.NewReflector()
	s := web.NewService(r)

	// Init API documentation schema.
	s.OpenAPISchema().SetTitle("Campaign Manager")
	s.OpenAPISchema().SetDescription("Campaign Manager Rest API.")
	s.OpenAPISchema().SetVersion("v0.1.0")

	// CORS
	s.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		//Enable Debugging for testing, consider disabling in production
		//Debug: true,
	}).Handler)

	{
		u := usecase.NewInteractor(func(ctx context.Context, input pingInput, output *pingOutput) error {
			output.Message = "OK"
			return nil
		})
		s.Get("/ping", u)
	}

	{
		u := usecase.NewInteractor(func(ctx context.Context, input lsfJobInput, output *lsfJobOutput) error {
			output.Jobs = lsf.Jobs()
			return nil
		})
		u.SetTitle("LSF Jobs")
		u.SetDescription("Returns all recent LSF Jobs.")
		s.Get("/lsf/job", u)
	}

	// generate opanapi file
	schema, err := r.Spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("docs/campaign-manager.yaml", schema, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// Swagger UI endpoint at /docs.
	s.Docs("/docs", swgui.New)

	// Start server.
	log.Println("http://localhost:8011/docs")
	if err := http.ListenAndServe("localhost:8011", s); err != nil {
		log.Fatal(err)
	}
}
