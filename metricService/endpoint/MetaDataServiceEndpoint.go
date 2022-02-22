package endpoint

import (
	"gin-operator/metricService/service"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func MakeMetadataServiceEndpoint(service service.IMetricService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return service.GetServiceName()
	}

}
