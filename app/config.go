package app

import (
	"github.com/spf13/viper"
)

type ConfigDatabase struct {
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
}

func SetDbConfiguration(envPath string) (c ConfigDatabase, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(envPath)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic("SETUP CONFIGURATION FAILED")
	}
	err = viper.Unmarshal(&c)
	return
}
