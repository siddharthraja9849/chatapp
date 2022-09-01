package utils

import (
	"chatapp/types"
	"fmt"
	"github.com/spf13/viper"
)

func LoadEnvConfigs(config *types.ENV) {
	viper.SetConfigFile(EnvFilePath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("%v+", err)
		panic("Unable to Read Configs")
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Errorf("%v+", err)
		panic("Unable to Marshal Configs")
	}
}
