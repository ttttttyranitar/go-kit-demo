package handler

import (
	"gin-operator/arithmatic/transport"
	"gin-operator/metricService/endpoint"
	"gin-operator/metricService/service"
	"github.com/go-kit/kit/transport/http"
)

func MetricServiceGetMetaHandler() *http.Server {
	var service service.MetricServiceImpl
	endpoint := endpoint.MakeMetadataServiceEndpoint(service)
	return http.NewServer(endpoint, transport.DecodeKubeServiceRequest, transport.EncodeKubeServiceResponse)
}
