package util

import "github.com/spf13/viper"

//Config stores all configuration of application
//the values are fetched using viper from config files or environment variable
type Config struct {
	DBdriver      string `mapstructure:"DB_DRVIER"`
	DBsoruce      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
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
	return
}
