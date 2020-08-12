package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
	"github.com/hanjin7278/gin-base-starter/logs"
	"github.com/spf13/viper"
)

var err error
var engin *gorose.Engin

func InitDB() {
	// 全局初始化数据库,并复用
	// 这里的engin需要全局保存,可以用全局变量,也可以用单例
	// 配置&gorose.Config{}是单一数据库配置
	// 如果配置读写分离集群,则使用&gorose.ConfigCluster{}
	// mysql示例, 记得导入mysql驱动 github.com/go-sql-driver/mysql
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		viper.GetString("mysql.username"), viper.GetString("mysql.password"), viper.GetString("mysql.host"), viper.GetString("mysql.port"),
		viper.GetString("mysql.database"))
	engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: url})

	if err != nil {
		logs.Error("数据库初始化出现错误", err)
		panic(err)
	}
}

func DB() gorose.IOrm {
	return engin.NewOrm()
}
