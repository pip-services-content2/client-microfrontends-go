package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-microfrontends-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type MicrofrontendsHttpCommandableClientV1Test struct {
	client  *version1.MicrofrontendsCommandableHttpClientV1
	fixture *MicrofrontendsClientFixtureV1
}

func newMicrofrontendsHttpCommandableClientV1Test() *MicrofrontendsHttpCommandableClientV1Test {
	return &MicrofrontendsHttpCommandableClientV1Test{}
}

func (c *MicrofrontendsHttpCommandableClientV1Test) setup(t *testing.T) *MicrofrontendsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewMicrofrontendsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewMicrofrontendsClientFixtureV1(c.client)

	return c.fixture
}

func (c *MicrofrontendsHttpCommandableClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestHttpCrudOperations(t *testing.T) {
	c := newMicrofrontendsHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
