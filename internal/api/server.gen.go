// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// File defines model for File.
type File struct {
	IsDir bool   `json:"isDir"`
	Name  string `json:"name"`
}

// DeleteFilesJSONBody defines parameters for DeleteFiles.
type DeleteFilesJSONBody struct {
	Path string `json:"path"`
}

// GetFilesJSONBody defines parameters for GetFiles.
type GetFilesJSONBody struct {
	Path string `json:"path"`
}

// DeleteFilesJSONRequestBody defines body for DeleteFiles for application/json ContentType.
type DeleteFilesJSONRequestBody DeleteFilesJSONBody

// GetFilesJSONRequestBody defines body for GetFiles for application/json ContentType.
type GetFilesJSONRequestBody GetFilesJSONBody

// PostFilesJSONRequestBody defines body for PostFiles for application/json ContentType.
type PostFilesJSONRequestBody = File

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Download plain/text file
	// (GET /file/{filePath})
	GetFileFilePath(w http.ResponseWriter, r *http.Request, filePath string)
	// Update file
	// (PUT /file/{filePath})
	PutFileFilePath(w http.ResponseWriter, r *http.Request, filePath string)
	// Delete files or folders
	// (DELETE /files)
	DeleteFiles(w http.ResponseWriter, r *http.Request)
	// List files including folders
	// (GET /files)
	GetFiles(w http.ResponseWriter, r *http.Request)
	// Add new file or folder
	// (POST /files)
	PostFiles(w http.ResponseWriter, r *http.Request)
	// Update files
	// (PUT /files)
	PutFiles(w http.ResponseWriter, r *http.Request)
	// Ping server
	// (GET /ping)
	GetPing(w http.ResponseWriter, r *http.Request)
	// List projects
	// (GET /projects)
	GetProjects(w http.ResponseWriter, r *http.Request)
	// Add new projects
	// (POST /projects)
	PostProjects(w http.ResponseWriter, r *http.Request)
	// List tasks
	// (GET /tasks)
	GetTasks(w http.ResponseWriter, r *http.Request)
	// Add new task
	// (POST /tasks)
	PostTasks(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Download plain/text file
// (GET /file/{filePath})
func (_ Unimplemented) GetFileFilePath(w http.ResponseWriter, r *http.Request, filePath string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update file
// (PUT /file/{filePath})
func (_ Unimplemented) PutFileFilePath(w http.ResponseWriter, r *http.Request, filePath string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete files or folders
// (DELETE /files)
func (_ Unimplemented) DeleteFiles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List files including folders
// (GET /files)
func (_ Unimplemented) GetFiles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add new file or folder
// (POST /files)
func (_ Unimplemented) PostFiles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update files
// (PUT /files)
func (_ Unimplemented) PutFiles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Ping server
// (GET /ping)
func (_ Unimplemented) GetPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List projects
// (GET /projects)
func (_ Unimplemented) GetProjects(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add new projects
// (POST /projects)
func (_ Unimplemented) PostProjects(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List tasks
// (GET /tasks)
func (_ Unimplemented) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add new task
// (POST /tasks)
func (_ Unimplemented) PostTasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetFileFilePath operation middleware
func (siw *ServerInterfaceWrapper) GetFileFilePath(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "filePath" -------------
	var filePath string

	err = runtime.BindStyledParameterWithOptions("simple", "filePath", chi.URLParam(r, "filePath"), &filePath, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filePath", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFileFilePath(w, r, filePath)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutFileFilePath operation middleware
func (siw *ServerInterfaceWrapper) PutFileFilePath(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "filePath" -------------
	var filePath string

	err = runtime.BindStyledParameterWithOptions("simple", "filePath", chi.URLParam(r, "filePath"), &filePath, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "filePath", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutFileFilePath(w, r, filePath)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteFiles operation middleware
func (siw *ServerInterfaceWrapper) DeleteFiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteFiles(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetFiles operation middleware
func (siw *ServerInterfaceWrapper) GetFiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFiles(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostFiles operation middleware
func (siw *ServerInterfaceWrapper) PostFiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostFiles(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutFiles operation middleware
func (siw *ServerInterfaceWrapper) PutFiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutFiles(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetPing operation middleware
func (siw *ServerInterfaceWrapper) GetPing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPing(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetProjects operation middleware
func (siw *ServerInterfaceWrapper) GetProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetProjects(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostProjects operation middleware
func (siw *ServerInterfaceWrapper) PostProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostProjects(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetTasks operation middleware
func (siw *ServerInterfaceWrapper) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTasks(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostTasks operation middleware
func (siw *ServerInterfaceWrapper) PostTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostTasks(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/file/{filePath}", wrapper.GetFileFilePath)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/file/{filePath}", wrapper.PutFileFilePath)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/files", wrapper.DeleteFiles)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/files", wrapper.GetFiles)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/files", wrapper.PostFiles)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/files", wrapper.PutFiles)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/ping", wrapper.GetPing)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/projects", wrapper.GetProjects)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/projects", wrapper.PostProjects)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/tasks", wrapper.GetTasks)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/tasks", wrapper.PostTasks)
	})

	return r
}

type GetFileFilePathRequestObject struct {
	FilePath string `json:"filePath"`
}

type GetFileFilePathResponseObject interface {
	VisitGetFileFilePathResponse(w http.ResponseWriter) error
}

type GetFileFilePath200PlaintextResponse struct {
	Body          io.Reader
	ContentLength int64
}

func (response GetFileFilePath200PlaintextResponse) VisitGetFileFilePathResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "plain/text")
	if response.ContentLength != 0 {
		w.Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	w.WriteHeader(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(w, response.Body)
	return err
}

type GetFileFilePath400Response struct {
}

func (response GetFileFilePath400Response) VisitGetFileFilePathResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetFileFilePath404Response struct {
}

func (response GetFileFilePath404Response) VisitGetFileFilePathResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PutFileFilePathRequestObject struct {
	FilePath string `json:"filePath"`
	Body     io.Reader
}

type PutFileFilePathResponseObject interface {
	VisitPutFileFilePathResponse(w http.ResponseWriter) error
}

type PutFileFilePath200Response struct {
}

func (response PutFileFilePath200Response) VisitPutFileFilePathResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PutFileFilePath400Response struct {
}

func (response PutFileFilePath400Response) VisitPutFileFilePathResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PutFileFilePath404Response struct {
}

func (response PutFileFilePath404Response) VisitPutFileFilePathResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type DeleteFilesRequestObject struct {
	Body *DeleteFilesJSONRequestBody
}

type DeleteFilesResponseObject interface {
	VisitDeleteFilesResponse(w http.ResponseWriter) error
}

type DeleteFiles204Response struct {
}

func (response DeleteFiles204Response) VisitDeleteFilesResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteFiles404Response struct {
}

func (response DeleteFiles404Response) VisitDeleteFilesResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetFilesRequestObject struct {
	Body *GetFilesJSONRequestBody
}

type GetFilesResponseObject interface {
	VisitGetFilesResponse(w http.ResponseWriter) error
}

type GetFiles200JSONResponse []File

func (response GetFiles200JSONResponse) VisitGetFilesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostFilesRequestObject struct {
	Body *PostFilesJSONRequestBody
}

type PostFilesResponseObject interface {
	VisitPostFilesResponse(w http.ResponseWriter) error
}

type PostFiles204Response struct {
}

func (response PostFiles204Response) VisitPostFilesResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type PutFilesRequestObject struct {
}

type PutFilesResponseObject interface {
	VisitPutFilesResponse(w http.ResponseWriter) error
}

type GetPingRequestObject struct {
}

type GetPingResponseObject interface {
	VisitGetPingResponse(w http.ResponseWriter) error
}

type GetPing200Response struct {
}

func (response GetPing200Response) VisitGetPingResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type GetProjectsRequestObject struct {
}

type GetProjectsResponseObject interface {
	VisitGetProjectsResponse(w http.ResponseWriter) error
}

type PostProjectsRequestObject struct {
}

type PostProjectsResponseObject interface {
	VisitPostProjectsResponse(w http.ResponseWriter) error
}

type GetTasksRequestObject struct {
}

type GetTasksResponseObject interface {
	VisitGetTasksResponse(w http.ResponseWriter) error
}

type PostTasksRequestObject struct {
}

type PostTasksResponseObject interface {
	VisitPostTasksResponse(w http.ResponseWriter) error
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Download plain/text file
	// (GET /file/{filePath})
	GetFileFilePath(ctx context.Context, request GetFileFilePathRequestObject) (GetFileFilePathResponseObject, error)
	// Update file
	// (PUT /file/{filePath})
	PutFileFilePath(ctx context.Context, request PutFileFilePathRequestObject) (PutFileFilePathResponseObject, error)
	// Delete files or folders
	// (DELETE /files)
	DeleteFiles(ctx context.Context, request DeleteFilesRequestObject) (DeleteFilesResponseObject, error)
	// List files including folders
	// (GET /files)
	GetFiles(ctx context.Context, request GetFilesRequestObject) (GetFilesResponseObject, error)
	// Add new file or folder
	// (POST /files)
	PostFiles(ctx context.Context, request PostFilesRequestObject) (PostFilesResponseObject, error)
	// Update files
	// (PUT /files)
	PutFiles(ctx context.Context, request PutFilesRequestObject) (PutFilesResponseObject, error)
	// Ping server
	// (GET /ping)
	GetPing(ctx context.Context, request GetPingRequestObject) (GetPingResponseObject, error)
	// List projects
	// (GET /projects)
	GetProjects(ctx context.Context, request GetProjectsRequestObject) (GetProjectsResponseObject, error)
	// Add new projects
	// (POST /projects)
	PostProjects(ctx context.Context, request PostProjectsRequestObject) (PostProjectsResponseObject, error)
	// List tasks
	// (GET /tasks)
	GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error)
	// Add new task
	// (POST /tasks)
	PostTasks(ctx context.Context, request PostTasksRequestObject) (PostTasksResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetFileFilePath operation middleware
func (sh *strictHandler) GetFileFilePath(w http.ResponseWriter, r *http.Request, filePath string) {
	var request GetFileFilePathRequestObject

	request.FilePath = filePath

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetFileFilePath(ctx, request.(GetFileFilePathRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetFileFilePath")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetFileFilePathResponseObject); ok {
		if err := validResponse.VisitGetFileFilePathResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PutFileFilePath operation middleware
func (sh *strictHandler) PutFileFilePath(w http.ResponseWriter, r *http.Request, filePath string) {
	var request PutFileFilePathRequestObject

	request.FilePath = filePath

	request.Body = r.Body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PutFileFilePath(ctx, request.(PutFileFilePathRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutFileFilePath")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PutFileFilePathResponseObject); ok {
		if err := validResponse.VisitPutFileFilePathResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteFiles operation middleware
func (sh *strictHandler) DeleteFiles(w http.ResponseWriter, r *http.Request) {
	var request DeleteFilesRequestObject

	var body DeleteFilesJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteFiles(ctx, request.(DeleteFilesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteFiles")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteFilesResponseObject); ok {
		if err := validResponse.VisitDeleteFilesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetFiles operation middleware
func (sh *strictHandler) GetFiles(w http.ResponseWriter, r *http.Request) {
	var request GetFilesRequestObject

	var body GetFilesJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetFiles(ctx, request.(GetFilesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetFiles")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetFilesResponseObject); ok {
		if err := validResponse.VisitGetFilesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostFiles operation middleware
func (sh *strictHandler) PostFiles(w http.ResponseWriter, r *http.Request) {
	var request PostFilesRequestObject

	var body PostFilesJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostFiles(ctx, request.(PostFilesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostFiles")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostFilesResponseObject); ok {
		if err := validResponse.VisitPostFilesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PutFiles operation middleware
func (sh *strictHandler) PutFiles(w http.ResponseWriter, r *http.Request) {
	var request PutFilesRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PutFiles(ctx, request.(PutFilesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutFiles")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PutFilesResponseObject); ok {
		if err := validResponse.VisitPutFilesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetPing operation middleware
func (sh *strictHandler) GetPing(w http.ResponseWriter, r *http.Request) {
	var request GetPingRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetPing(ctx, request.(GetPingRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPing")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetPingResponseObject); ok {
		if err := validResponse.VisitGetPingResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetProjects operation middleware
func (sh *strictHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	var request GetProjectsRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetProjects(ctx, request.(GetProjectsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProjects")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetProjectsResponseObject); ok {
		if err := validResponse.VisitGetProjectsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostProjects operation middleware
func (sh *strictHandler) PostProjects(w http.ResponseWriter, r *http.Request) {
	var request PostProjectsRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostProjects(ctx, request.(PostProjectsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostProjects")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostProjectsResponseObject); ok {
		if err := validResponse.VisitPostProjectsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetTasks operation middleware
func (sh *strictHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	var request GetTasksRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasks(ctx, request.(GetTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasks")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetTasksResponseObject); ok {
		if err := validResponse.VisitGetTasksResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostTasks operation middleware
func (sh *strictHandler) PostTasks(w http.ResponseWriter, r *http.Request) {
	var request PostTasksRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostTasks(ctx, request.(PostTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTasks")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostTasksResponseObject); ok {
		if err := validResponse.VisitPostTasksResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xVy27bOhD9FWLuXaqW2mSlXdrARYC2MPpYBVkw0thmKpEsOUpiGPr3YkjZsh1FbpKN",
	"N5YhDg/PYzhaQ2FqazRq8pCvwRdLrGX4O1UV8tM6Y9GRwvBW+Uvl+A+tLEIOt8ZUKDW0CWhZ486KJ6f0",
	"Ato2AYd/GuWwhPw6ViUdzk3Ly0rPDW8s0RdOWVJGQw6fZG2lWmjxVWq5QCe+oydxMbuaQAKkiNk9KYIE",
	"7tH5iJBN3k8yZmYsamkV5HA2ySZnkICVtAx60rmqMF3z70zSsuV3CyR+sGzJZK5KyOEzEjsy7QoDhpM1",
	"EjoP+fUaFB9p41K0AuZ9cW8BuQaTzmg+Bh9lbYOYlHenZAInFnlg5A3DeGu0j1l8yDJ+FEYT6kDZVlLp",
	"lPCR+jCHImmTA7N/NEWB3rNZ5xF1f/2bET22YH6CFaEnLOOm86FN+h0+Kk9KL0Rwhg/2TV1Lt4IcLs2D",
	"rowsD6EZ0DYDGcya08ggCP9oytVr7W+HkzyBTH7ZUhJ2MbRJvB8+4lRI+DSUy/B+GsrGzJHWVqoI+9I7",
	"z2R2LdqfMoHY0VESqm62GZnbOyxo2N7zZ+0VD4qWQhuxofrqbg4+BOe8ME7MTVVyX7bJ6EA5YdOyF5FR",
	"hHXY+L/DOeTwX9p/WtLuu5JOu+vdnS6dk6vRebRn8RflqTNY6aJqSo5ix2dr/NDUMP7NTh/X9Nau2xN6",
	"UZZC40O81NteOjYXo8CewXMXO9qaWm6Pkc/djNdfMKjEpvJADOMIj+6eFYSDneGm86OHb2pGJIV22IKN",
	"5v8vcBvTe0QmS9L/HmX6MxQcoxlhRjkeBdoQZKzQcH8DAAD//4pgaq6+CQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
