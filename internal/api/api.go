package api

//func (s *CampaignManagerService) Ping(
//	ctx context.Context,
//	req *connect.Request[protocmv1.PingRequest],
//) (*connect.Response[protocmv1.PingResponse], error) {
//	log.Println("Request headers: ", req.Header())
//	res := connect.NewResponse(&protocmv1.PingResponse{
//		Message: "OK",
//	})
//	return res, nil
//}

//func (s *CampaignManagerService) NewProject(
//	ctx context.Context,
//	req *connect.Request[protocmv1.NewProjectRequest],
//) (*connect.Response[protocmv1.NewProjectResponse], error) {
//	log.Println("Request headers: ", req.Header())
//	log.Println("Request Msg: ", req.Msg)
//
//	projectId := uint64(11)
//
//	res := connect.NewResponse(&protocmv1.NewProjectResponse{
//		ProjectId: &projectId,
//	})
//	return res, nil
//}
//
//func (s *CampaignManagerService) GetSettings(
//	ctx context.Context,
//	req *connect.Request[protocmv1.GetSettingsRequest],
//) (*connect.Response[protocmv1.GetSettingsResponse], error) {
//
//	settings, err := cmdb.GetSettings(s.db)
//	if err != nil {
//		return nil, err
//	}
//
//	if settings == nil {
//		res := connect.NewResponse(&protocmv1.GetSettingsResponse{
//			WorkingDir:  "",
//			LsfUsername: "",
//			LsfPassword: "",
//		})
//
//		return res, nil
//	}
//
//	res := connect.NewResponse(&protocmv1.GetSettingsResponse{
//		WorkingDir:  settings.WorkingDir(),
//		LsfUsername: settings.LSFUsername(),
//		LsfPassword: settings.LSFPassword(),
//	})
//	return res, nil
//}
//
//func (s *CampaignManagerService) SetSettings(
//	_ context.Context,
//	req *connect.Request[protocmv1.SetSettingsRequest],
//) (*connect.Response[protocmv1.SetSettingsResponse], error) {
//
//	settings := types.NewSettings(req.Msg.WorkingDir, req.Msg.LsfUsername, req.Msg.LsfPassword)
//	err := cmdb.SetSettings(s.db, settings)
//	if err != nil {
//		return nil, err
//	}
//
//	res := connect.NewResponse(&protocmv1.SetSettingsResponse{})
//	return res, nil
//}

//	{
//		u := usecase.NewInteractor(func(ctx context.Context, input pingInput, output *pingOutput) error {
//			output.Message = "OK"
//			return nil
//		})
//		s.Get("/ping", u)
//		u.SetTitle("Ping server")
//		u.SetDescription("Ping the server to test the connection.")
//
//	}
//
//	{
//		u := usecase.NewInteractor(func(ctx context.Context, input lsfJobInput, output *lsfJobOutput) error {
//			output.Jobs = lsf.Jobs()
//			return nil
//		})
//		u.SetTitle("LSF Jobs")
//		u.SetDescription("Returns all recent LSF Jobs.")
//		s.Get("/lsf/job", u)
//	}
//
//	{
//		u := usecase.NewInteractor(func(ctx context.Context, input struct{}, output *settingsOutput) error {
//			settings, err := cmdb.GetSettings(s.db)
//			if err != nil {
//				return nil, err
//			}
//			return nil
//		})
//		//u.SetTitle("LSF Jobs")
//		//u.SetDescription("Returns all recent LSF Jobs.")
//		s.Get("/settings", u)
//	}

//func Server() {
//	c := config.NewConfig(true)
//
//	data := db.ConnectDB()
//	defer db.DisconnectDB(data)
//
//	// openapi
//	oapi := openapi3.Reflector{}
//	oapi.Spec = &openapi3.Spec{Openapi: "3.0.3"}
//	oapi.Spec.Info.
//		WithTitle("Campaign Manager").
//		WithVersion("0.1.0").
//		WithDescription("Campaign Manager Rest API.")
//
//	r := chi.NewRouter()
//	r.Use(middleware.Logger)
//	r.Use(cors.Handler(cors.Options{
//
//		// TODO set this based on env variable
//		AllowedOrigins: []string{"http://localhost:5173"},
//
//		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
//		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
//		ExposedHeaders:   []string{"Link"},
//		AllowCredentials: false,
//		MaxAge:           300,
//	}))
//
//	{
//		resp := "OK"
//		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
//			_, err := w.Write([]byte(resp))
//			if err != nil {
//				slog.Error("Error in get /ping response write", err)
//			}
//		})
//
//		getOp, err := oapi.NewOperationContext(http.MethodGet, "/ping")
//		if err != nil {
//			slog.Error("oapi new op err", err)
//		}
//		getOp.AddRespStructure(resp)
//		err = oapi.AddOperation(getOp)
//		if err != nil {
//			slog.Error("oapi add op err", err)
//		}
//	}
//
//	if c.Dev() {
//		schema, err := oapi.Spec.MarshalYAML()
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		err = os.WriteFile("./internal/api/openapi.yaml", schema, 0666)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		r.Get("/api/openapi.json", func(w http.ResponseWriter, r *http.Request) {
//			w.Header().Set("Content-Type", "application/json")
//			err := json.NewEncoder(w).Encode(oapi.Spec)
//			if err != nil {
//				if c.Dev() {
//					http.Error(w, "Internal error: "+err.Error(), http.StatusInternalServerError)
//					return
//				}
//
//				http.Error(w, "Internal error", http.StatusInternalServerError)
//				return
//			}
//		})
//
//		r.Mount("/api/docs", v5emb.New(
//			oapi.Spec.Title(),
//			"/api/openapi.json",
//			"/api/docs/",
//		))
//		println("docs at http://localhost:3000/api/docs")
//	}
//
//	if err := http.ListenAndServe("localhost:3000", r); err != nil {
//		log.Fatal(err)
//	}
//}
