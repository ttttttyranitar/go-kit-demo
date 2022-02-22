package viper

import (
	log "github.com/sirupsen/logrus"
	_ "github.com/spf13/viper"
	viper2 "github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	remoteConfig *viper2.Viper
)

func SetViper() {

	viper2.SetConfigType("yaml")
	viper2.AddConfigPath(".")
	viper2.SetConfigName("service-dev")
	err := viper2.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"cause": err,
		}).Error("init yaml error")
	}
	log.WithFields(log.Fields{
		"info": "viper loaded config file",
	}).Info("config  file has loaded")

}
