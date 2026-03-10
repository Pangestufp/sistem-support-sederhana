package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	DBUsername  string `mapstructure:"DB_USERNAME"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBUrl       string `mapstructure:"DB_URL"`
	RedisUrl    string `mapstructure:"REDIS_URL"`
	DBDatabase  string `mapstructure:"DB_DATABASE"`
	SecretKey   string `mapstructure:"SECRET_KEY"`
	FrontendURL string `mapstructure:"FRONTEND_URL"`
}

var ENV Config

func LoadConfig() {
	viper.AutomaticEnv()

	viper.BindEnv("PORT")
	viper.BindEnv("DB_URL")
	viper.BindEnv("DB_USERNAME")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_DATABASE")
	viper.BindEnv("REDIS_URL")
	viper.BindEnv("SECRET_KEY")
	viper.BindEnv("FRONTEND_URL")

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal("Cannot unmarshal config:", err)
	}
}
