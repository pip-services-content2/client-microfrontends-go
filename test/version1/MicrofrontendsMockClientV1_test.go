package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-microfrontends-go/version1"
)

type MicrofrontendsMockClientV1 struct {
	client  *version1.MicrofrontendsMockClientV1
	fixture *MicrofrontendsClientFixtureV1
}

func newMicrofrontendsMockClientV1() *MicrofrontendsMockClientV1 {
	return &MicrofrontendsMockClientV1{}
}

func (c *MicrofrontendsMockClientV1) setup(t *testing.T) *MicrofrontendsClientFixtureV1 {
	c.client = version1.NewMicrofrontendsMockClientV1()
	c.fixture = NewMicrofrontendsClientFixtureV1(c.client)

	return c.fixture
}

func (c *MicrofrontendsMockClientV1) teardown(t *testing.T) {
	c.client = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newMicrofrontendsMockClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
