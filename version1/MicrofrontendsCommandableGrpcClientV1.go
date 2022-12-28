package version1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type MicrofrontendsCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewMicrofrontendsCommandableGrpcClientV1() *MicrofrontendsCommandableGrpcClientV1 {
	return &MicrofrontendsCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/microfrontends"),
	}
}

func (c *MicrofrontendsCommandableGrpcClientV1) GetMicrofrontends(ctx context.Context, correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (result cdata.DataPage[*MicrofrontendV1], err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_microfrontends", correlationId, params)
	if err != nil {
		return *cdata.NewEmptyDataPage[*MicrofrontendV1](), err
	}

	return clients.HandleHttpResponse[cdata.DataPage[*MicrofrontendV1]](res, correlationId)
}

func (c *MicrofrontendsCommandableGrpcClientV1) GetMicrofrontendById(ctx context.Context, correlationId string, microfrontendId string) (result *MicrofrontendV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"microfrontend_id", microfrontendId,
	)

	res, err := c.CallCommand(ctx, "get_microfrontend_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*MicrofrontendV1](res, correlationId)
}

func (c *MicrofrontendsCommandableGrpcClientV1) CreateMicrofrontend(ctx context.Context, correlationId string, microfrontend *MicrofrontendV1) (result *MicrofrontendV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"microfrontend", microfrontend,
	)

	res, err := c.CallCommand(ctx, "create_microfrontend", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*MicrofrontendV1](res, correlationId)
}

func (c *MicrofrontendsCommandableGrpcClientV1) UpdateMicrofrontend(ctx context.Context, correlationId string, microfrontend *MicrofrontendV1) (result *MicrofrontendV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"microfrontend", microfrontend,
	)

	res, err := c.CallCommand(ctx, "update_microfrontend", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*MicrofrontendV1](res, correlationId)
}

func (c *MicrofrontendsCommandableGrpcClientV1) DeleteMicrofrontendById(ctx context.Context, correlationId string, microfrontendId string) (result *MicrofrontendV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"microfrontend_id", microfrontendId,
	)

	res, err := c.CallCommand(ctx, "delete_microfrontend_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*MicrofrontendV1](res, correlationId)
}
