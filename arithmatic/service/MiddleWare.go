package service

import (
	"gin-operator/pb"
	log "github.com/go-kit/kit/log"
)

type logMiddleware struct {
	logger log.Logger
	next   IKubeService
}

func (l logMiddleware) Add(param1 int64, param2 int64) (*pb.Result, error) {
	l.logger.Log("开始计算加法")
	return l.next.Add(param1, param2)
}

func (l logMiddleware) Mul(param1 int64, param2 int64) (*pb.Result, error) {
	l.logger.Log("开始计算乘法")
	return l.next.Mul(param1, param2)
}

type Middleware func(service IKubeService) IKubeService

func LogMiddleware(logger log.Logger) Middleware {
	return func(next IKubeService) IKubeService {
		return &logMiddleware{logger, next}
	}
}
