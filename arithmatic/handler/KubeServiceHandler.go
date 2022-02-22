package handler

import (
	"gin-operator/arithmatic/endpoints"
	"gin-operator/arithmatic/service"
	"gin-operator/arithmatic/transport"
	"github.com/go-kit/kit/transport/http"
)

func KubeAddServiceHandler() *http.Server {
	kubeService := service.KubeService()
	serviceEndpoint := endpoints.MakeKubeAddServiceEndpoint(kubeService)
	//返回一个服务对象
	return http.NewServer(serviceEndpoint, transport.DecodeKubeServiceRequest, transport.EncodeKubeServiceResponse)
}

func KubeMultiServiceHandler() *http.Server {
	KubeService := service.KubeService()
	serviceEndpoint := endpoints.MakeKubeMultiServiceEndpoint(KubeService)
	return http.NewServer(serviceEndpoint, transport.DecodeKubeServiceRequest, transport.EncodeKubeServiceResponse)
}
