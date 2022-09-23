package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:host`
	Port string `mapstructure:port`
}

type RedisConfig struct {
	Address string `mapstructure:address`
	Password string `mapstructure:password`
	DB int `mapstructure:db`
}

type Config struct {
	Server ServerConfig `mapstructure:server`
	Redis RedisConfig `mapstructure:redis`
}

var vp *viper.Viper


func LoadConfig(path string) (Config, error) {
	vp = viper.New()
	var config Config

	vp.AddConfigPath(path)
	vp.SetConfigName("cfg")

	err := vp.ReadInConfig()
	if err != nil {
		log.Println(err)
		return Config{}, err
	}
	err = vp.Unmarshal(&config)
	if err != nil {
		log.Println(err)
		return Config{}, err
	}

	return config, nil
}