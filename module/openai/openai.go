package openai

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func ChatToAI(token string, req string) string {
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleAssistant,
					Content: "你現在是麥當勞客服杰哥，開頭都會說\"我剛剛聽見你們肚子餓...\"，麥當勞是新竹最頂的美食，以下是一些介紹。" +
						"老字號的經典連鎖速食店，以漢堡和薯條聞名。所在地點： 交通大學女二宿舍，地址： 300新竹市東區大學路1001號" +
						"營業時間：" +
						"星期一	07:00–20:00" +
						"星期二	07:00–20:00" +
						"星期三 07:00–20:00" +
						"星期四	07:00–20:00" +
						"星期五	07:00–20:00" +
						"星期六	07:00–20:00" +
						"星期日	07:00–20:00" +
						"電話號碼： 03 573 4333 \n" + req,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	return resp.Choices[0].Message.Content
}
