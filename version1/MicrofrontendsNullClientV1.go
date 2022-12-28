package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type MicrofrontendsNullClientV1 struct {
}

func NewMicrofrontendsNullClientV1() *MicrofrontendsNullClientV1 {
	return &MicrofrontendsNullClientV1{}
}

func (c *MicrofrontendsNullClientV1) GetMicrofrontends(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*MicrofrontendV1], err error) {
	return *data.NewEmptyDataPage[*MicrofrontendV1](), nil
}

func (c *MicrofrontendsNullClientV1) GetMicrofrontendById(ctx context.Context, correlationId string, microfrontendId string) (result *MicrofrontendV1, err error) {
	return nil, nil
}

func (c *MicrofrontendsNullClientV1) CreateMicrofrontend(ctx context.Context, correlationId string, microfrontend *MicrofrontendV1) (result *MicrofrontendV1, err error) {
	return microfrontend, nil
}

func (c *MicrofrontendsNullClientV1) UpdateMicrofrontend(ctx context.Context, correlationId string, microfrontend *MicrofrontendV1) (result *MicrofrontendV1, err error) {
	return microfrontend, nil
}

func (c *MicrofrontendsNullClientV1) DeleteMicrofrontendById(ctx context.Context, correlationId string, microfrontendId string) (result *MicrofrontendV1, err error) {
	return nil, nil
}
