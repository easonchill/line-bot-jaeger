package router

import (
	"line-bot-jaeger/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/callback", controller.Callback)
	r.POST("/sendMassage", controller.SendMassage)

	r.GET("/messages", controller.GetAllMessage)
	r.GET("/message/:userid", controller.GetUserMessage)
}
