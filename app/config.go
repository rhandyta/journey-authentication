package app

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigDatabase struct {
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
}

func SetDbConfiguration() (c ConfigDatabase, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	fmt.Println(err)
	if err != nil {
		panic("SETUP CONFIGURATION FAILED")
	}
	err = viper.Unmarshal(&c)
	return
}
