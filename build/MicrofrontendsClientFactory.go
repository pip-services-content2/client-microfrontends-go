package build

import (
	clients1 "github.com/pip-services-content2/client-microfrontends-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type MicrofrontendsClientFactory struct {
	*cbuild.Factory
}

func NewMicrofrontendsClientFactory() *MicrofrontendsClientFactory {
	c := &MicrofrontendsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-microfrontends", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-microfrontends", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-microfrontends", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-microfrontends", "client", "grpc", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("service-microfrontends", "client", "commandable-grpc", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewMicrofrontendsNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewMicrofrontendsMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewMicrofrontendsCommandableHttpClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewMicrofrontendsGrpcClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewMicrofrontendsCommandableGrpcClientV1)

	return c
}
