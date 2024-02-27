package controller

import (
	"fmt"
	"net/http"

	"line-bot-jaeger/model"

	"github.com/gin-gonic/gin"
)

func GetAllMessage(c *gin.Context) {
	allMessage := model.GetAllMessage(c.Request.Context())

	wrapResponse(c, http.StatusOK, "Success", allMessage)
	return
}

func GetUserMessage(c *gin.Context) {
	userid := c.Param("userid")
	fmt.Println(userid)
	allMessage := model.GetUserMessage(c.Request.Context(), userid)

	wrapResponse(c, http.StatusOK, "Success", allMessage)
	return
}
