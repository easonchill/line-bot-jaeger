package controller

import (
	"errors"
	"time"

	"line-bot-jaeger/config"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data        any    `json:"data"`
	ResponseMsg any    `json:"responseMsg"`
	Datetime    string `json:"datetime"`
}

func wrapResponse(c *gin.Context, statusCode int, RspMsg, data any) {
	timeZone, _ := time.LoadLocation("Asia/Taipei")

	c.JSON(statusCode, Response{
		Data:        data,
		ResponseMsg: RspMsg,
		Datetime:    time.Now().In(timeZone).Format(time.RFC3339)})

	return
}

func getLineConfig(c *gin.Context) (*config.Line, error) {
	lineConfig, exists := c.MustGet("Line").(config.Line)
	if !exists {
		return nil, errors.New("line config load error")
	}

	return &lineConfig, nil
}
