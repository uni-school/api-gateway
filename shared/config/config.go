package config

import (
	"errors"
	"flag"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/uni-school/api-gateaway/shared/constant"
)

var (
	AppConfig   *DefaultConfig
	Environment constant.EnvironmentType
)

func ConfigApps() {
	var (
		environment = flag.String("env", "", "input the environment type")
	)

	flag.Parse()

	switch constant.EnvironmentType(*environment) {
	case constant.DEV:
		viper.SetConfigFile("./environments/dev.application.yaml")
	case constant.STAG:
		viper.SetConfigFile("./environments/stag.application.yaml")
	case constant.PROD:
		viper.SetConfigFile("./environments/prod.application.yaml")
	case constant.TEST:
		viper.SetConfigFile("./environments/test.application.yaml")
	default:
		panic(errors.New("input environment type [ dev | stag | prod | test]"))
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	var conf DefaultConfig
	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Panic(err)
	}

	Environment = constant.EnvironmentType(*environment)
	AppConfig = &conf
}
