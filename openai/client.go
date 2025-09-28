package openai

import (
    "context"
    "log"
    "os"
	"fmt"

    "github.com/joho/godotenv"
    openai "github.com/sashabaranov/go-openai"
)

func NewClient() *openai.Client {
	fmt.Println("Entering function")
    _ = godotenv.Load()

    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        log.Fatal("Missing OPENAI_API_KEY. Make sure it's in your .env file.")
    }

    return openai.NewClient(apiKey)
}


func AskGPT(input string) (string, error){
	client := NewClient();

	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
            Model: openai.GPT5, 
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleSystem,
                    Content: `You are a helpful assistant that converts natural language into safe Linux shell commands with explanations. 
                    Your responses will first simply list out the command for the user, then it will provide an explanation`,
                },
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: input,
                },
            },
		},
	)

	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}