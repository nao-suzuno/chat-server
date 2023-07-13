package main

import (
	"encoding/json"
	gpt "chat-server/gpt"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
	
	"golang.org/x/net/context"
	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
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

	openai := gpt.NewGpt(os.Getenv("TOKEN"), "gpt-3.5-turbo")
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
