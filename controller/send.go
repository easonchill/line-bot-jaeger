package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type SendMessageReq struct {
	UserID string `json:"userid" form:"userid"`
	Text   string `json:"text" form:"text"`
}

func SendMassage(c *gin.Context) {
	lineConfig, err := getLineConfig(c)
	if err != nil {
		wrapResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	bot, err := messaging_api.NewMessagingApiAPI(lineConfig.Token)
	if err != nil {
		wrapResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	req := SendMessageReq{}
	if err = c.ShouldBind(&req); err != nil {
		wrapResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rsp, err := bot.PushMessage(
		&messaging_api.PushMessageRequest{
			To: req.UserID,
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: req.Text,
				},
			},
		}, "",
	)
	if err != nil {
		wrapResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	wrapResponse(c, http.StatusOK, "Success", rsp)
	return
}
