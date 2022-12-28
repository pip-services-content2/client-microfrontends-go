package version1

import (
	"context"

	"github.com/pip-services-content2/client-microfrontends-go/protos"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type MicrofrontendsGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewMicrofrontendsGrpcClientV1() *MicrofrontendsGrpcClientV1 {
	return &MicrofrontendsGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("microfrontends_v1.Microfrontends"),
	}
}

func (c *MicrofrontendsGrpcClientV1) GetMicrofrontends(ctx context.Context, correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (result cdata.DataPage[*MicrofrontendV1], err error) {
	timing := c.Instrument(ctx, correlationId, "microfrontends_v1.get_microfrontends")
	defer timing.EndTiming(ctx, err)

	req := &protos.MicrofrontendPageRequest{
		CorrelationId: correlationId,
	}

	if filter != nil {
		req.Filter = filter.Value()
	}

	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.MicrofrontendPageReply)
	err = c.CallWithContext(ctx, "get_microfrontends", correlationId, req, reply)
	if err != nil {
		return *cdata.NewEmptyDataPage[*MicrofrontendV1](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *cdata.NewEmptyDataPage[*MicrofrontendV1](), err
	}

	result = toMicrofrontendPage(reply.Page)

	return result, nil
}

func (c *MicrofrontendsGrpcClientV1) GetMicrofrontendById(ctx context.Context, correlationId string, microfrontendId string) (result *MicrofrontendV1, err error) {
	timing := c.Instrument(ctx, correlationId, "microfrontends_v1.get_microfrontend_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.MicrofrontendIdRequest{
		CorrelationId:   correlationId,
		MicrofrontendId: microfrontendId,
	}

	reply := new(protos.MicrofrontendObjectReply)
	err = c.CallWithContext(ctx, "get_microfrontend_by_id", correlationId, req, reply)

	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toMicrofrontend(reply.Microfrontend)

	return result, nil
}

func (c *MicrofrontendsGrpcClientV1) CreateMicrofrontend(ctx context.Context, correlationId string, microfrontend *MicrofrontendV1) (result *MicrofrontendV1, err error) {
	timing := c.Instrument(ctx, correlationId, "microfrontends_v1.create_microfrontend")
	defer timing.EndTiming(ctx, err)

	req := &protos.MicrofrontendObjectRequest{
		CorrelationId: correlationId,
		Microfrontend: fromMicrofrontend(microfrontend),
	}

	reply := new(protos.MicrofrontendObjectReply)
	err = c.CallWithContext(ctx, "create_microfrontend", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toMicrofrontend(reply.Microfrontend)

	return result, nil
}

func (c *MicrofrontendsGrpcClientV1) UpdateMicrofrontend(ctx context.Context, correlationId string, microfrontend *MicrofrontendV1) (result *MicrofrontendV1, err error) {
	timing := c.Instrument(ctx, correlationId, "microfrontends_v1.update_microfrontend")
	defer timing.EndTiming(ctx, err)

	req := &protos.MicrofrontendObjectRequest{
		CorrelationId: correlationId,
		Microfrontend: fromMicrofrontend(microfrontend),
	}

	reply := new(protos.MicrofrontendObjectReply)
	err = c.CallWithContext(ctx, "update_microfrontend", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toMicrofrontend(reply.Microfrontend)

	return result, nil
}

func (c *MicrofrontendsGrpcClientV1) DeleteMicrofrontendById(ctx context.Context, correlationId string, microfrontendId string) (result *MicrofrontendV1, err error) {
	timing := c.Instrument(ctx, correlationId, "microfrontends_v1.delete_microfrontend_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.MicrofrontendIdRequest{
		CorrelationId:   correlationId,
		MicrofrontendId: microfrontendId,
	}

	reply := new(protos.MicrofrontendObjectReply)
	err = c.CallWithContext(ctx, "delete_microfrontend_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toMicrofrontend(reply.Microfrontend)

	return result, nil
}
