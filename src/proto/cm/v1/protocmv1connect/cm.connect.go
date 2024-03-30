// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/cm/v1/cm.proto

package protocmv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/campaign-manager/src/proto/cm/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PingServiceName is the fully-qualified name of the PingService service.
	PingServiceName = "proto.cm.v1.PingService"
	// NewProjectServiceName is the fully-qualified name of the NewProjectService service.
	NewProjectServiceName = "proto.cm.v1.NewProjectService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PingServicePingProcedure is the fully-qualified name of the PingService's Ping RPC.
	PingServicePingProcedure = "/proto.cm.v1.PingService/Ping"
	// NewProjectServiceNewProjectProcedure is the fully-qualified name of the NewProjectService's
	// NewProject RPC.
	NewProjectServiceNewProjectProcedure = "/proto.cm.v1.NewProjectService/NewProject"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	pingServiceServiceDescriptor                = v1.File_proto_cm_v1_cm_proto.Services().ByName("PingService")
	pingServicePingMethodDescriptor             = pingServiceServiceDescriptor.Methods().ByName("Ping")
	newProjectServiceServiceDescriptor          = v1.File_proto_cm_v1_cm_proto.Services().ByName("NewProjectService")
	newProjectServiceNewProjectMethodDescriptor = newProjectServiceServiceDescriptor.Methods().ByName("NewProject")
)

// PingServiceClient is a client for the proto.cm.v1.PingService service.
type PingServiceClient interface {
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
}

// NewPingServiceClient constructs a client for the proto.cm.v1.PingService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPingServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &pingServiceClient{
		ping: connect.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+PingServicePingProcedure,
			connect.WithSchema(pingServicePingMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// pingServiceClient implements PingServiceClient.
type pingServiceClient struct {
	ping *connect.Client[v1.PingRequest, v1.PingResponse]
}

// Ping calls proto.cm.v1.PingService.Ping.
func (c *pingServiceClient) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// PingServiceHandler is an implementation of the proto.cm.v1.PingService service.
type PingServiceHandler interface {
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
}

// NewPingServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPingServiceHandler(svc PingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	pingServicePingHandler := connect.NewUnaryHandler(
		PingServicePingProcedure,
		svc.Ping,
		connect.WithSchema(pingServicePingMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.cm.v1.PingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PingServicePingProcedure:
			pingServicePingHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPingServiceHandler struct{}

func (UnimplementedPingServiceHandler) Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.cm.v1.PingService.Ping is not implemented"))
}

// NewProjectServiceClient is a client for the proto.cm.v1.NewProjectService service.
type NewProjectServiceClient interface {
	NewProject(context.Context, *connect.Request[v1.NewProjectRequest]) (*connect.Response[v1.NewProjectResponse], error)
}

// NewNewProjectServiceClient constructs a client for the proto.cm.v1.NewProjectService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNewProjectServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) NewProjectServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &newProjectServiceClient{
		newProject: connect.NewClient[v1.NewProjectRequest, v1.NewProjectResponse](
			httpClient,
			baseURL+NewProjectServiceNewProjectProcedure,
			connect.WithSchema(newProjectServiceNewProjectMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// newProjectServiceClient implements NewProjectServiceClient.
type newProjectServiceClient struct {
	newProject *connect.Client[v1.NewProjectRequest, v1.NewProjectResponse]
}

// NewProject calls proto.cm.v1.NewProjectService.NewProject.
func (c *newProjectServiceClient) NewProject(ctx context.Context, req *connect.Request[v1.NewProjectRequest]) (*connect.Response[v1.NewProjectResponse], error) {
	return c.newProject.CallUnary(ctx, req)
}

// NewProjectServiceHandler is an implementation of the proto.cm.v1.NewProjectService service.
type NewProjectServiceHandler interface {
	NewProject(context.Context, *connect.Request[v1.NewProjectRequest]) (*connect.Response[v1.NewProjectResponse], error)
}

// NewNewProjectServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNewProjectServiceHandler(svc NewProjectServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	newProjectServiceNewProjectHandler := connect.NewUnaryHandler(
		NewProjectServiceNewProjectProcedure,
		svc.NewProject,
		connect.WithSchema(newProjectServiceNewProjectMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.cm.v1.NewProjectService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case NewProjectServiceNewProjectProcedure:
			newProjectServiceNewProjectHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedNewProjectServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedNewProjectServiceHandler struct{}

func (UnimplementedNewProjectServiceHandler) NewProject(context.Context, *connect.Request[v1.NewProjectRequest]) (*connect.Response[v1.NewProjectResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.cm.v1.NewProjectService.NewProject is not implemented"))
}