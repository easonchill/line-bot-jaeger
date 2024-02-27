package main

import (
	"github.com/gin-gonic/gin"
	"line-bot-jaeger/config"
	"line-bot-jaeger/model"
	"line-bot-jaeger/router"
	"log"
)

func main() {

	cfg := config.NewConfig()

	mongoClient := model.InitMongoDB(&cfg.MongoDB)
	model.InitCollection(mongoClient, cfg.MongoDB.Database)

	r := gin.Default()

	r.Use(func(c *gin.Context) {

		c.Set("Line", cfg.Line)
		c.Next()
	})

	router.SetupRouter(r)

	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Println(err)
	}

}
