package config

import (
	"fmt"
	"github.com/hanjin7278/gin-base-starter/logs"
	"github.com/spf13/viper"
	"os"
)

/**
读取配置文件
*/
func InitReadConfig() {
	var env string
	if len(os.Args) > 1 {
		env = os.Args[1]
	} else {
		env = "dev"
	}
	configName := fmt.Sprintf("app-%s", env)
	viper.SetConfigName(configName)
	viper.SetConfigType("yml")
	viper.AddConfigPath("./res")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logs.Error("读取配置文件错误: \n", err)
		panic(fmt.Errorf("Fatal errors config file: %s \n", err))
	}
	logs.Info("配置文件读取完成...")
}
