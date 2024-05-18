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

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// FileTreeEntry defines model for FileTreeEntry.
type FileTreeEntry struct {
	IsDir bool   `json:"isDir"`
	Path  string `json:"path"`
}

// FileTreePath defines model for FileTreePath.
type FileTreePath struct {
	Path string `json:"path"`
}

// PostFileTreeJSONRequestBody defines body for PostFileTree for application/json ContentType.
type PostFileTreeJSONRequestBody = FileTreeEntry

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Download plain/text file
	// (GET /file/{filePath})
	GetFileFilePath(w http.ResponseWriter, r *http.Request, filePath string)
	// Update file
	// (PUT /file/{filePath})
	PutFileFilePath(w http.ResponseWriter, r *http.Request, filePath string)
	// List files and folders in root (/)
	// (GET /fileTree)
	GetFileTree(w http.ResponseWriter, r *http.Request)
	// Add new file or folder
	// (POST /fileTree)
	PostFileTree(w http.ResponseWriter, r *http.Request)
	// Delete files or folders
	// (DELETE /fileTree/{path})
	DeleteFileTreePath(w http.ResponseWriter, r *http.Request, path string)
	// List files and folders in the given path
	// (GET /fileTree/{path})
	GetFileTreePath(w http.ResponseWriter, r *http.Request, path string)
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

// List files and folders in root (/)
// (GET /fileTree)
func (_ Unimplemented) GetFileTree(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add new file or folder
// (POST /fileTree)
func (_ Unimplemented) PostFileTree(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete files or folders
// (DELETE /fileTree/{path})
func (_ Unimplemented) DeleteFileTreePath(w http.ResponseWriter, r *http.Request, path string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List files and folders in the given path
// (GET /fileTree/{path})
func (_ Unimplemented) GetFileTreePath(w http.ResponseWriter, r *http.Request, path string) {
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

// GetFileTree operation middleware
func (siw *ServerInterfaceWrapper) GetFileTree(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFileTree(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostFileTree operation middleware
func (siw *ServerInterfaceWrapper) PostFileTree(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostFileTree(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteFileTreePath operation middleware
func (siw *ServerInterfaceWrapper) DeleteFileTreePath(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "path" -------------
	var path string

	err = runtime.BindStyledParameterWithOptions("simple", "path", chi.URLParam(r, "path"), &path, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "path", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteFileTreePath(w, r, path)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetFileTreePath operation middleware
func (siw *ServerInterfaceWrapper) GetFileTreePath(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "path" -------------
	var path string

	err = runtime.BindStyledParameterWithOptions("simple", "path", chi.URLParam(r, "path"), &path, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "path", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFileTreePath(w, r, path)
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
		r.Get(options.BaseURL+"/fileTree", wrapper.GetFileTree)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/fileTree", wrapper.PostFileTree)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/fileTree/{path}", wrapper.DeleteFileTreePath)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/fileTree/{path}", wrapper.GetFileTreePath)
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

type GetFileTreeRequestObject struct {
}

type GetFileTreeResponseObject interface {
	VisitGetFileTreeResponse(w http.ResponseWriter) error
}

type GetFileTree200JSONResponse []FileTreeEntry

func (response GetFileTree200JSONResponse) VisitGetFileTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostFileTreeRequestObject struct {
	Body *PostFileTreeJSONRequestBody
}

type PostFileTreeResponseObject interface {
	VisitPostFileTreeResponse(w http.ResponseWriter) error
}

type PostFileTree201JSONResponse FileTreePath

func (response PostFileTree201JSONResponse) VisitPostFileTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostFileTree409JSONResponse Error

func (response PostFileTree409JSONResponse) VisitPostFileTreeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

type DeleteFileTreePathRequestObject struct {
	Path string `json:"path"`
}

type DeleteFileTreePathResponseObject interface {
	VisitDeleteFileTreePathResponse(w http.ResponseWriter) error
}

type DeleteFileTreePath204Response struct {
}

func (response DeleteFileTreePath204Response) VisitDeleteFileTreePathResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteFileTreePath404Response struct {
}

func (response DeleteFileTreePath404Response) VisitDeleteFileTreePathResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetFileTreePathRequestObject struct {
	Path string `json:"path"`
}

type GetFileTreePathResponseObject interface {
	VisitGetFileTreePathResponse(w http.ResponseWriter) error
}

type GetFileTreePath200JSONResponse []FileTreeEntry

func (response GetFileTreePath200JSONResponse) VisitGetFileTreePathResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
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
	// List files and folders in root (/)
	// (GET /fileTree)
	GetFileTree(ctx context.Context, request GetFileTreeRequestObject) (GetFileTreeResponseObject, error)
	// Add new file or folder
	// (POST /fileTree)
	PostFileTree(ctx context.Context, request PostFileTreeRequestObject) (PostFileTreeResponseObject, error)
	// Delete files or folders
	// (DELETE /fileTree/{path})
	DeleteFileTreePath(ctx context.Context, request DeleteFileTreePathRequestObject) (DeleteFileTreePathResponseObject, error)
	// List files and folders in the given path
	// (GET /fileTree/{path})
	GetFileTreePath(ctx context.Context, request GetFileTreePathRequestObject) (GetFileTreePathResponseObject, error)
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

// GetFileTree operation middleware
func (sh *strictHandler) GetFileTree(w http.ResponseWriter, r *http.Request) {
	var request GetFileTreeRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetFileTree(ctx, request.(GetFileTreeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetFileTree")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetFileTreeResponseObject); ok {
		if err := validResponse.VisitGetFileTreeResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostFileTree operation middleware
func (sh *strictHandler) PostFileTree(w http.ResponseWriter, r *http.Request) {
	var request PostFileTreeRequestObject

	var body PostFileTreeJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostFileTree(ctx, request.(PostFileTreeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostFileTree")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostFileTreeResponseObject); ok {
		if err := validResponse.VisitPostFileTreeResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteFileTreePath operation middleware
func (sh *strictHandler) DeleteFileTreePath(w http.ResponseWriter, r *http.Request, path string) {
	var request DeleteFileTreePathRequestObject

	request.Path = path

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteFileTreePath(ctx, request.(DeleteFileTreePathRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteFileTreePath")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteFileTreePathResponseObject); ok {
		if err := validResponse.VisitDeleteFileTreePathResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetFileTreePath operation middleware
func (sh *strictHandler) GetFileTreePath(w http.ResponseWriter, r *http.Request, path string) {
	var request GetFileTreePathRequestObject

	request.Path = path

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetFileTreePath(ctx, request.(GetFileTreePathRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetFileTreePath")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetFileTreePathResponseObject); ok {
		if err := validResponse.VisitGetFileTreePathResponse(w); err != nil {
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

	"H4sIAAAAAAAC/9RWT08bPxD9Kv7NrweQttlQuHRvtECF1FZRS0+Ig7ueJKa7tjt2gCja716Nd/Nnk00C",
	"CKn0Qlb2+M2bN89jZpDb0lmDJnjIZuDzMZYyfp4TWeIPR9YhBY1xuUTv5Qj5M0wdQgY+kDYjqKoECH9P",
	"NKGC7HoReFMlcKELvCLEcxNouomp/ZmmFcSf1hYoDVQJOBnGvIMPsnQFb/JKGmw61AVCsodEPJ40CVaZ",
	"DBrcNpGXyHZT8bI2Q8tICn1O2gVtDWTwUZZO6pERX6SRIyTxDX0Qp4PLHmPrEFOuB0ECd0i+Ruj3jnp9",
	"FsY6NNJpyOC41+8dQy1VrCKSTWf8l8useG2EgX+4WMlkLhVk8AkD63HRBEYMkiUGJA/Z9Qy0aSSABIws",
	"md1wGbwsPdAEk8Y8bf3SVQHFwYQKgSa3SpuRMIgK1X+HHbreMLp31vi6Me/6ff7JrQloYiWukNqkAR/C",
	"0rddpqyStR58n+Q5es8antSo7f2vViyxRaTNhaIPqOpDJ12HzFt80D5wXVEwTuwnZSnZ8HBm701hpVqH",
	"jhafdLRmMHnVrYl6fLBq+tyuVN0NfgWt+uGUDNh0p0rq28QjY981ijF7fSudK3QeD6e3nhmt6qQDlvHg",
	"G8IhZPB/upzPaTOc0/Y0rRYtkkRyutPyrUI/a1+L5oU0SgxtoZC80EaQtUEcpIfRndZ32dP6ds3bDLG7",
	"3CdU2WWZoxdPNqj9sKEgr4ucUAZUvdpa718sef3SbssqC0KppiJ61vfWmniqlDB4X5vfUtPFtnHTmZs/",
	"AwoLDLjZz7O43lLhMRPH/YWH4GTrnBD3OoyFsWLelWdP66hGczcWosY5tG8AvGrp/sVZFMYoRvoOzbJR",
	"qeMCd0zjAe8/4YER88g1VowjPNLd/EY5sreY1/8qb00+j2kT2Cx3AbZzyj4Gbj4ElohMNkj/ayfTqxiw",
	"j2YNs5PjXqA5QcaKo/xPAAAA///pDu01eAwAAA==",
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
