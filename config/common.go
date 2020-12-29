package config

import (
	"github.com/tkanos/gonfig"
)

type CommonConfiguration struct {
	ServerPort int
	DBServer   string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	PrivateKey string
}

var GlobalConfig CommonConfiguration

func LoadCommonConfig() (conf CommonConfiguration, err error) {
	err = gonfig.GetConf("common.config.json", &conf)
	GlobalConfig = conf
	return conf, err
}
