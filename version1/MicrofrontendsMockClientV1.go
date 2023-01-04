package version1

import (
	"context"
	"strings"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type MicrofrontendsMockClientV1 struct {
	microfrontends []*MicrofrontendV1
}

func NewMicrofrontendsMockClientV1() *MicrofrontendsMockClientV1 {
	return &MicrofrontendsMockClientV1{
		microfrontends: make([]*MicrofrontendV1, 0),
	}
}

func (c *MicrofrontendsMockClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(search))
}

func (c *MicrofrontendsMockClientV1) matchSearch(item *MicrofrontendV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchString(item.Id, search) {
		return true
	}
	if c.matchString(item.Name, search) {
		return true
	}
	if c.matchString(item.Description, search) {
		return true
	}
	return false
}

func (c *MicrofrontendsMockClientV1) composeFilter(filter *data.FilterParams) func(item *MicrofrontendV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	search, searchOk := filter.GetAsNullableString("search")
	id, idOk := filter.GetAsNullableString("id")
	mType, mTypeOk := filter.GetAsNullableString("type")
	exposed, exposedOk := filter.GetAsNullableString("exposed_module")

	return func(item *MicrofrontendV1) bool {
		if searchOk && !c.matchSearch(item, search) {
			return false
		}
		if idOk && id != item.Id {
			return false
		}
		if mTypeOk && mType != item.Type {
			return false
		}
		if exposedOk && exposed != item.ExposedModule {
			return false
		}

		return true
	}
}

func (c *MicrofrontendsMockClientV1) GetMicrofrontends(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*MicrofrontendV1], err error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*MicrofrontendV1, 0)
	for _, v := range c.microfrontends {
		item := *v
		if filterFunc(&item) {
			items = append(items, &item)
		}
	}
	return *data.NewDataPage(items, len(c.microfrontends)), nil
}

func (c *MicrofrontendsMockClientV1) GetMicrofrontendById(ctx context.Context, correlationId string, microfrontendId string) (result *MicrofrontendV1, err error) {
	for _, v := range c.microfrontends {
		if v.Id == microfrontendId {
			buf := *v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *MicrofrontendsMockClientV1) CreateMicrofrontend(ctx context.Context, correlationId string, microfrontend *MicrofrontendV1) (result *MicrofrontendV1, err error) {
	buf := *microfrontend
	c.microfrontends = append(c.microfrontends, &buf)
	return microfrontend, nil
}

func (c *MicrofrontendsMockClientV1) UpdateMicrofrontend(ctx context.Context, correlationId string, microfrontend *MicrofrontendV1) (result *MicrofrontendV1, err error) {
	if microfrontend == nil {
		return nil, nil
	}

	var index = -1
	for i, v := range c.microfrontends {
		if v.Id == microfrontend.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	buf := *microfrontend
	c.microfrontends[index] = &buf
	return microfrontend, nil
}

func (c *MicrofrontendsMockClientV1) DeleteMicrofrontendById(ctx context.Context, correlationId string, microfrontendId string) (result *MicrofrontendV1, err error) {
	var index = -1
	for i, v := range c.microfrontends {
		if v.Id == microfrontendId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}
	var item = c.microfrontends[index]
	if index < len(c.microfrontends) {
		c.microfrontends = append(c.microfrontends[:index], c.microfrontends[index+1:]...)
	} else {
		c.microfrontends = c.microfrontends[:index]
	}
	return item, nil
}
