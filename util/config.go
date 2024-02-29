package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER" default:"postgres"`
	DBSource      string `mapstructure:"DB_URL"`
	ServerAddress string `mapstructure:"API_URL" default:"0.0.0.0:8080"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	{
		return
	}
}
