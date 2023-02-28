package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type OpenAIClient struct {
	Organization string
	APIKey       string
	RateLimit    int64
}

func (o OpenAIClient) RequestCompletion(body CompletionRequestBody) (*CompletionResponse, error) {
	client := &http.Client{}

	byts, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(byts)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", reqBody)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if o.Organization != "" {
		req.Header.Add("OpenAI-Organization", o.Organization)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+o.APIKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var resp CompletionResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func main() {
	// Load env
	LoadConfig()

	newAiClient := OpenAIClient{
		Organization: os.Getenv("OPENAI_ORG"),
		APIKey:       os.Getenv("OPENAI_API_KEY"),
	}

	model := flag.String("m", "text-davinci-003", "language model to use. Default: text-davinci-003")
	maxTokens := flag.Int("t", 2048, "Max tokens to use. Default: 2048")
	flag.Parse()

	prompt := os.Args[1]

	newReq := CompletionRequestBody{
		Model:     *model,
		MaxTokens: int64(*maxTokens),
		Prompt:    prompt,
	}

	response, err := newAiClient.RequestCompletion(newReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your response: %v\n", response)

}

func LoadConfig() {
	// TODO: move this
	err := godotenv.Load("/etc/env/openai.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
