package config_manager

import "github.com/gin-gonic/gin"

func router() *gin.Engine {
	r := gin.Default()
	r.GET("", handleListConfigs)
	return r
}

func handleListConfigs(c *gin.Context) {

}
