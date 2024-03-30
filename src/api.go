package cm

import (
	"connectrpc.com/connect"
	"context"
	_ "github.com/campaign-manager/src/proto/cm/v1"
	protocmv1 "github.com/campaign-manager/src/proto/cm/v1"
	"github.com/campaign-manager/src/proto/cm/v1/protocmv1connect"
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

func Server() {
	pinger := &PingServer{}
	mux := http.NewServeMux()
	path, handler := protocmv1connect.NewPingServiceHandler(pinger)
	mux.Handle(path, handler)
	err := http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
