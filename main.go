package main

import (
	"log"
	"net/http"

	"line-bot-jaeger/config"
	"line-bot-jaeger/model"
	"line-bot-jaeger/router"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
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

	r.Use(func(c *gin.Context) {
		c.Set("OpenAI", cfg.OpenAI.Token)
		c.Next()
	})

	router.SetupRouter(r)

	log.Fatal(http.Serve(autocert.NewListener("api.easonchill.dev"), r))
}
