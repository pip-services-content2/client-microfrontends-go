package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IMicrofrontendsClientV1 interface {
	GetMicrofrontends(ctx context.Context, correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result data.DataPage[*MicrofrontendV1], err error)

	GetMicrofrontendById(ctx context.Context, correlationId string, microfrontendId string) (result *MicrofrontendV1, err error)

	CreateMicrofrontend(ctx context.Context, correlationId string, microfrontend *MicrofrontendV1) (result *MicrofrontendV1, err error)

	UpdateMicrofrontend(ctx context.Context, correlationId string, microfrontend *MicrofrontendV1) (result *MicrofrontendV1, err error)

	DeleteMicrofrontendById(ctx context.Context, correlationId string, microfrontendId string) (result *MicrofrontendV1, err error)
}
