package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/hanjin7278/gin-base-starter/boot"
	"github.com/hanjin7278/gin-base-starter/logs"
	"github.com/hanjin7278/gin-base-starter/routers"
	"github.com/spf13/viper"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {

	runMode := viper.GetString("server.runMode")
	gin.SetMode(runMode)
	r := routers.InitRouter()

	endPoint := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	logs.Info("[info] start http server listening", endPoint)
	r.Run(endPoint)
}
