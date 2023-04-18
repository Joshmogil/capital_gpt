package ai

import (
	"context"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

type gpt4Connection struct {
	apiToken string
	openaiOrg string
	// client openai.Client
}


func (conn *gpt4Connection) getResponse(prompt string) string {
	client := openai.NewClient(conn.apiToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	
	

	if err != nil {
		log.Fatal(err.Error())
	}

	return resp.Choices[0].Message.Content
}



