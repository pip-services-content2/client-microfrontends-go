package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-microfrontends-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type microfrontendsGrpcClientV1Test struct {
	client  *version1.MicrofrontendsGrpcClientV1
	fixture *MicrofrontendsClientFixtureV1
}

func newMicrofrontendsGrpcClientV1Test() *microfrontendsGrpcClientV1Test {
	return &microfrontendsGrpcClientV1Test{}
}

func (c *microfrontendsGrpcClientV1Test) setup(t *testing.T) *MicrofrontendsClientFixtureV1 {
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

func (c *microfrontendsGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGrpcCrudOperations(t *testing.T) {
	c := newMicrofrontendsGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
