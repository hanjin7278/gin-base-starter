package boot

import (
	"github.com/hanjin7278/gin-base-starter/config"
	"github.com/hanjin7278/gin-base-starter/logs"
	_ "github.com/hanjin7278/gin-base-starter/utils/errors"
)

/**
初始化引导配置相关
*/
func init() {
	logs.Info("初始化引导配置...[starting]")
	//读取配置文件信息
	config.InitReadConfig()
	//初始化数据库连接信息
	config.InitDB()
	logs.Info("初始化引导配置...[ending]")
}
