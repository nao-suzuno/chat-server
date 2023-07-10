package gpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type IGpt interface {
	GenerateText(message []Chat) (string, error)
}

type Gpt struct {
	apiKey string
	model  string
}

func NewGpt(apiKey string, model string) IGpt {
	return &Gpt{
		apiKey: apiKey,
		model:  model,
	}
}

type Chat struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (gpt *Gpt) GenerateText(messages []Chat) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"


	jsonData, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "", errors.New("error marshaling JSON")
	}

	jsonString := string(jsonData)
	jsonStr := []byte(fmt.Sprintf(`{
		"model": "%s",
		"messages": %s
	}`, gpt.model, jsonString))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "error before sending a request to api", err
	}
	if gpt.apiKey == "" {
		return "invalid_api_key", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+gpt.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "error after sending a request to api", err
	}
	defer resp.Body.Close()

	var m map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {

		return "error after sending a request to api", err
	}
	if errVal, ok := m["error"].(map[string]interface{}); ok {

		errorCode := errVal["code"].(string)

		return errorCode, errors.New(errorCode)
	}

	content := m["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return content, err
}

