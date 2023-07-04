package config

import "fmt"
import "github.com/spf13/viper"

var config Config

func init() {
	// 设置配置文件名和路径
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("解析配置文件失败: %s", err))
	}
}

func GetConfig() *Config {
	return &config
}

type Config struct {
	MySQL struct {
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	} `mapstructure:"mysql"`
	URL struct {
		Ratings string `mapstructure:"ratings"`
	} `mapstructure:"url"`
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}
