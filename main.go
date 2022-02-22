package main

import (
	"gin-operator/config/viper"
	"gin-operator/routers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	viper.SetViper()

	//logrus组件打印日志
	log.WithFields(log.Fields{
		"port": 9090,
	}).Info("service Starting")

	http.ListenAndServe("localhost:9090", routers.Router())
}
