package api

import "github.com/spf13/viper"

type Config struct {
	BASE_URL_TREASURY       string `mapstructure:"BASE_URL_TREASURY"`
	RATES_EXCHANGE_ENDPOINT string `mapstructure:"RATES_EXCHANGE_ENDPOINT"`
	SERVER_HOST             string `mapstructure:"SERVER_HOST"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
