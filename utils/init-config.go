package utils

import "github.com/spf13/viper"

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigType("yml")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
