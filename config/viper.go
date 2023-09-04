package config

import (
	"log"

	"github.com/spf13/viper"
)

type EnvVars struct {
	Port         string `mapstructure:"PORT"`
	KafkaAddress string `mapstructure:"KAFKA_ADDRESS"`
	KafkaPort    string `mapstructure:"KAFKA_PORT"`
	KafkaTopic   string `mapstructure:"KAFKA_TOPIC"`
	WsPort       string `mapstructure:"WS_PORT"`
}

func LoadConfig() (config EnvVars, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	return
}
