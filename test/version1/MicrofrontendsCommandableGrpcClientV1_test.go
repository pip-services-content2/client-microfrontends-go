package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-microfrontends-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type microfrontendsGrpcCommandableClientV1Test struct {
	client  *version1.MicrofrontendsGrpcClientV1
	fixture *MicrofrontendsClientFixtureV1
}

func newMicrofrontendsGrpcCommandableClientV1Test() *microfrontendsGrpcCommandableClientV1Test {
	return &microfrontendsGrpcCommandableClientV1Test{}
}

func (c *microfrontendsGrpcCommandableClientV1Test) setup(t *testing.T) *MicrofrontendsClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewMicrofrontendsGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewMicrofrontendsClientFixtureV1(c.client)

	return c.fixture
}

func (c *microfrontendsGrpcCommandableClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableGrpcCrudOperations(t *testing.T) {
	c := newMicrofrontendsGrpcCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
