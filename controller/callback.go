package controller

import (
	"fmt"
	"log"
	"net/http"

	"line-bot-jaeger/model"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func Callback(c *gin.Context) {
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

	cb, err := webhook.ParseRequest(lineConfig.ChannelSecret, c.Request)
	//若訊息解析失敗或驗證就返回錯誤
	if err != nil {
		log.Printf("Cannot parse request: %+v\n", err)
		if err.Error() == linebot.ErrInvalidSignature.Error() {
			wrapResponse(c, http.StatusBadRequest, linebot.ErrInvalidSignature.Error(), nil)
		} else {
			wrapResponse(c, http.StatusInternalServerError, linebot.ErrInvalidSignature.Error(), nil)
		}
		return
	}

	for _, event := range cb.Events {
		log.Printf("/callback called%+v...\n", event)

		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				_, err := model.InsertMessage(c.Request.Context(), model.NewMessage(e.Source.(webhook.UserSource).UserId, message.Text, e.Timestamp))
				if err != nil {
					wrapResponse(c, http.StatusInternalServerError, err.Error(), nil)
				}

				if _, err = bot.ReplyMessage(
					&messaging_api.ReplyMessageRequest{
						ReplyToken: e.ReplyToken,
						Messages: []messaging_api.MessageInterface{
							messaging_api.TextMessage{
								Text: message.Text,
							},
						},
					},
				); err != nil {
					log.Print(err)
				} else {
					wrapResponse(c, 200, "Success", "echo success")
				}
			case webhook.StickerMessageContent:
				replyMessage := fmt.Sprintf(
					"sticker id is %s, stickerResourceType is %s", message.StickerId, message.StickerResourceType)
				if _, err = bot.ReplyMessage(
					&messaging_api.ReplyMessageRequest{
						ReplyToken: e.ReplyToken,
						Messages: []messaging_api.MessageInterface{
							messaging_api.TextMessage{
								Text: replyMessage,
							},
						},
					}); err != nil {
					log.Print(err)
				} else {
					log.Println("Sent sticker reply.")
				}
			default:
				log.Printf("Unsupported message content: %T\n", e.Message)
			}
		default:
			log.Printf("Unsupported message: %T\n", event)
		}
	}
}
