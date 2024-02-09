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
	if err := viper.ReadInConfig(); err != nil {
		// Handle error, e.g., return an error instead of panicking
		return ConfigDatabase{}, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		// Handle error, e.g., return an error instead of panicking
		return ConfigDatabase{}, err
	}
	return c, nil
}
