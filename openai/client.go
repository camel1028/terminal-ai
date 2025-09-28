package openai

import (
    "context"
    "log"
    "os"
	"fmt"
    "strings"

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


// func AskGPT(input string) (string, error){
// 	client := NewClient();

// 	response, err := client.CreateChatCompletion(
// 		context.Background(),
// 		openai.ChatCompletionRequest{
//             Model: openai.GPT5, 
//             Messages: []openai.ChatCompletionMessage{
//                 {
//                     Role:    openai.ChatMessageRoleSystem,
//                     Content: `You are a helpful assistant that converts natural language into safe Linux shell commands with explanations. 
//                     Your responses will first simply list out the command for the user, then it will provide an explanation`,
//                 },
//                 {
//                     Role:    openai.ChatMessageRoleUser,
//                     Content: input,
//                 },
//             },
// 		},
// 	)

// 	if err != nil {
// 		return "", err
// 	}

// 	return response.Choices[0].Message.Content, nil
// }
func AskGPT(prompt string) (string, error) {
	client := NewClient()

	req := openai.ChatCompletionRequest{
		Model:  openai.GPT4,
		Stream: true, 
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: `You are a helpful assistant that converts natural language into safe Linux shell commands with explanations.`,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	stream, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("failed to start stream: %w", err)
	}
	defer stream.Close()

	fmt.Println("Loading...")

	var fullResponse strings.Builder

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}

		token := resp.Choices[0].Delta.Content
		fmt.Print(token)                      // listing out tokens
		fullResponse.WriteString(token)   
	}

	fmt.Println()
	return fullResponse.String(), nil
}