package endpoints

import (
	"context"
	"gin-operator/arithmatic/service"
	"gin-operator/pb"
	"github.com/go-kit/kit/endpoint"
)

//构造函数，返回一个Endpoint函数，里面实际参与运算的时service层的函数
//在handler中需要调用此函数生成server对象。
func MakeKubeAddServiceEndpoint(kubeService service.IKubeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(pb.CalculateParams)
		result, err := kubeService.Add(req.Param1, req.Param2)
		return result, err
	}
}

func MakeKubeMultiServiceEndpoint(kubeService service.IKubeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(pb.CalculateParams)
		result, err := kubeService.Mul(req.Param1, req.Param2)
		return result, err
	}
}
