package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"

	"github.com/pip-services-content2/client-microfrontends-go/version1"
	"github.com/stretchr/testify/assert"
)

type MicrofrontendsClientFixtureV1 struct {
	Client version1.IMicrofrontendsClientV1

	MICROFRONTEND1 *version1.MicrofrontendV1
	MICROFRONTEND2 *version1.MicrofrontendV1
}

func NewMicrofrontendsClientFixtureV1(client version1.IMicrofrontendsClientV1) *MicrofrontendsClientFixtureV1 {
	return &MicrofrontendsClientFixtureV1{
		Client: client,
		MICROFRONTEND1: &version1.MicrofrontendV1{
			Id:            "1",
			Name:          "Microfrontend 1",
			Description:   "Main module",
			PathPrefix:    "md1",
			Icon:          "icon1",
			Type:          "vue",
			RemoteEntry:   "/remote",
			ExposedModule: "module1",
			ElementName:   "main_module",
			Params:        make(map[string]string, 0),
		},
		MICROFRONTEND2: &version1.MicrofrontendV1{
			Id:            "2",
			Name:          "Microfrontend 2",
			Description:   "Second module",
			PathPrefix:    "md2",
			Icon:          "icon2",
			Type:          "vue",
			RemoteEntry:   "/remote",
			ExposedModule: "module2",
			ElementName:   "second_module",
			Params:        make(map[string]string, 0),
		},
	}
}

func (c *MicrofrontendsClientFixtureV1) clear() {
	page, _ := c.Client.GetMicrofrontends(context.Background(), "", data.NewEmptyFilterParams(), data.NewEmptyPagingParams())

	for _, microfrontend := range page.Data {
		c.Client.DeleteMicrofrontendById(context.Background(), "", microfrontend.Id)
	}
}

func (c *MicrofrontendsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one microfrontend
	microfrontend, err := c.Client.CreateMicrofrontend(context.Background(), "", c.MICROFRONTEND1)
	assert.Nil(t, err)

	assert.NotNil(t, microfrontend)
	assert.Equal(t, microfrontend.Name, c.MICROFRONTEND1.Name)
	assert.Equal(t, microfrontend.Description, c.MICROFRONTEND1.Description)
	assert.Equal(t, microfrontend.PathPrefix, c.MICROFRONTEND1.PathPrefix)
	assert.Equal(t, microfrontend.Icon, c.MICROFRONTEND1.Icon)
	assert.Equal(t, microfrontend.Type, c.MICROFRONTEND1.Type)
	assert.Equal(t, microfrontend.RemoteEntry, c.MICROFRONTEND1.RemoteEntry)
	assert.Equal(t, microfrontend.ExposedModule, c.MICROFRONTEND1.ExposedModule)
	assert.Equal(t, microfrontend.ElementName, c.MICROFRONTEND1.ElementName)

	microfrontend1 := microfrontend

	// Create another microfrontend
	microfrontend, err = c.Client.CreateMicrofrontend(context.Background(), "", c.MICROFRONTEND2)
	assert.Nil(t, err)

	assert.NotNil(t, microfrontend)
	assert.Equal(t, microfrontend.Name, c.MICROFRONTEND2.Name)
	assert.Equal(t, microfrontend.Description, c.MICROFRONTEND2.Description)
	assert.Equal(t, microfrontend.PathPrefix, c.MICROFRONTEND2.PathPrefix)
	assert.Equal(t, microfrontend.Icon, c.MICROFRONTEND2.Icon)
	assert.Equal(t, microfrontend.Type, c.MICROFRONTEND2.Type)
	assert.Equal(t, microfrontend.RemoteEntry, c.MICROFRONTEND2.RemoteEntry)
	assert.Equal(t, microfrontend.ExposedModule, c.MICROFRONTEND2.ExposedModule)
	assert.Equal(t, microfrontend.ElementName, c.MICROFRONTEND2.ElementName)

	// Get all accounts
	page, err1 := c.Client.GetMicrofrontends(context.Background(), "", data.NewEmptyFilterParams(), data.NewEmptyPagingParams())
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 2)

	// Update the microfrontend
	microfrontend1.Name = "Updated Name 1"
	microfrontend, err = c.Client.UpdateMicrofrontend(context.Background(), "", microfrontend1)
	assert.Nil(t, err)

	assert.NotNil(t, microfrontend)
	assert.Equal(t, microfrontend.Name, "Updated Name 1")
	assert.Equal(t, microfrontend.Id, microfrontend1.Id)

	microfrontend1 = microfrontend

	// Delete microfrontend
	_, err = c.Client.DeleteMicrofrontendById(context.Background(), "", microfrontend1.Id)
	assert.Nil(t, err)

	// Try to get delete microfrontend
	microfrontend, err = c.Client.GetMicrofrontendById(context.Background(), "", microfrontend1.Id)
	assert.Nil(t, err)

	assert.Nil(t, microfrontend)
}
