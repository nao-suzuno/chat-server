package main

import (
	"encoding/json"
	gpt "chat-server/gpt"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Result string `json:"result"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Content-Type":                    "application/json",
		"Access-Control-Allow-Origin":     "*",
		"Access-Control-Allow-Methods":    "OPTIONS,POST,GET",
		"Access-Control-Allow-Headers":    "Origin,Authorization,Accept,X-Requested-With",
		"Access-Control-Allow-Credential": "true",
	}

	text := request.QueryStringParameters["text"]

	openai := gpt.NewGpt("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", "gpt-3.5-turbo")
	chat := []gpt.Chat{
		{
			Role:    "user",
			Content: text,
		},
	}
	result, _ := openai.GenerateText(chat)

	response_json := Response{
		Result: result,
  }
  jsonBytes, _ := json.Marshal(response_json)

	return events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:      string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
    lambda.Start(handler)
}
