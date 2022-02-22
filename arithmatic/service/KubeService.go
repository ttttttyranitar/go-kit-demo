package service

import "gin-operator/pb"

/**
 *KubeService接口定义
 */
type IKubeService interface {
	Add(param1 int64, param2 int64) (*pb.Result, error)
	Mul(param1 int64, param2 int64) (*pb.Result, error)
}

/**
接口实现结构体
*/
type KubeServiceImpl struct {
}

//对于包内定义的对象（绑定了方法的结构体），我们需要提供构造方法
func KubeService() *KubeServiceImpl {
	return &KubeServiceImpl{}
}

/**
*为KubeServiceImpl绑定方法
 */
func (s *KubeServiceImpl) Add(param1 int64, param2 int64) (*pb.Result, error) {
	res := param1 + param2
	return &pb.Result{Result: res}, nil
}

func (s *KubeServiceImpl) Mul(param1 int64, param2 int64) (*pb.Result, error) {
	res := param1 * param2
	return &pb.Result{Result: res}, nil
}
