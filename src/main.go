package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"gin-rest-api/src/app/routes"
)

func main() {
	viper.AddConfigPath("./src/conf")
	viper.SetConfigName("appsettings")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	router := gin.Default()
	routes.RegisterRoutes(router)

	server := fmt.Sprintf(":%s", viper.Get("server.port"))
	router.Run(server)
}
