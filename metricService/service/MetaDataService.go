package service

import (
	"gin-operator/pb"
	"github.com/spf13/viper"
)

const (
	SERVICE_NAME string = "app.name"
	SERVICE_PORT string = "app.port"
)

type IMetricService interface {
	GetServiceName() (*pb.Metric, error)
}

type MetricServiceImpl struct {
}

//获得服务yml中描述的服务名和端口信息。
func (receiver MetricServiceImpl) GetServiceName() (*pb.Metric, error) {

	name := viper.GetString(SERVICE_NAME)
	port := viper.GetInt32(SERVICE_PORT)
	return &pb.Metric{
		ServiceName: name,
		Port:        port,
	}, nil

}
