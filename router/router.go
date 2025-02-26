package router

import (
	"backend/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Setup() {
	r := gin.New()
	r.Use(log.GinLogger(zap.L()), log.GinRecovery(zap.L(), true))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run(":8080")
}
