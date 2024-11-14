package core

import (
	"fmt"
	"github.com/spf13/viper"
)

var defaultConfig = "./configs/conf.yaml"

func InitConfig(configFile string) (err error) {
	if configFile == "" {
		configFile = defaultConfig
	}
	viper.SetConfigFile(configFile)
	// 配置文件格式
	viper.SetConfigType("yaml")
	// 读取配置文件内容
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Sprintf("viper read config error: %v", err))
		return
	}
	// 反序列化配置参数到全局变量 GlobalConfig 中
	if err = viper.Unmarshal(&GlobalConfig); err != nil {
		fmt.Println("config unmarshal error:", err)
		return
	}
	return nil
}
