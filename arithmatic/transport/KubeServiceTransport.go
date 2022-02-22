package transport

import (
	"encoding/json"
	"gin-operator/pb"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

//go-kit将请求参数提取的模块独立出来，这个函数将从请求中获取参数，必要时组装到DTO之中。
func DecodeKubeServiceRequest(ctx context.Context, req *http.Request) (request interface{}, err error) {
	param1, _ := strconv.Atoi(req.FormValue("param1"))
	param2, _ := strconv.Atoi(req.FormValue("param2"))
	//将参数封装到请求对象中返回
	return pb.CalculateParams{Param1: int64(param1), Param2: int64(param2)}, nil

}

//接口执行结果封装至response中，这一步是必要的。
func EncodeKubeServiceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
