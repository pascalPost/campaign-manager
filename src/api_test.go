package cm_test

import (
	"connectrpc.com/connect"
	"context"
	cm "github.com/campaign-manager/src"
	protocmv1 "github.com/campaign-manager/src/proto/cm/v1"
	"github.com/campaign-manager/src/proto/cm/v1/protocmv1connect"
	"github.com/stretchr/testify/assert"
	"go.akshayshah.org/memhttp/memhttptest"
	"testing"
)

func TestPing(t *testing.T) {
	_, handle := protocmv1connect.NewPingServiceHandler(&cm.PingServer{})

	srv := memhttptest.New(t, handle)

	client := protocmv1connect.NewPingServiceClient(
		srv.Client(),
		srv.URL(),
		connect.WithHTTPGet(),
	)

	req := connect.NewRequest(&protocmv1.PingRequest{})
	res, err := client.Ping(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "OK", res.Msg.Message)
}
