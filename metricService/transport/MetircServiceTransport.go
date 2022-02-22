package transport

import (
	"encoding/json"
	"golang.org/x/net/context"
	"net/http"
)

//封装response header 与response body 中的内容，这一操作是必要的
func EncodeKubeServiceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
