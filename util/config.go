package util

import "github.com/spf13/viper"

// Config 保存从配置文件中读取的变量
type Config struct {
	DbSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig 读取path的配置文件内容
func LoadConfig(path string) (config Config, err error) {
	// 设置配置文件位置、名称、类型   json xml env
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// 自动覆盖
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
