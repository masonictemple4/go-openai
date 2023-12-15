package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

// TODO: Move the Client object over here.
type OpenAIClient struct {
	Organization string
	APIKey       string
	RateLimit    int64

	mu sync.Mutex

	taskQueue chan func() error
	logger    *log.Logger

	http.Client
}

const (
	OpenAIRoleSystem    = "system"
	OpenAIRoleUser      = "user"
	OpenAIRoleAssistant = "assistant"
	OpenAIRoleTool      = "tool"
	OpenAIRoleFunction  = "function"
)

func ValidOpenAIRole(role string) bool {
	switch role {
	case OpenAIRoleSystem, OpenAIRoleUser, OpenAIRoleAssistant, OpenAIRoleTool, OpenAIRoleFunction:
		return true
	default:
		return false
	}
}

type OpenAIOpt struct {
}

func New(opts *[]OpenAIOpt) *OpenAIClient {
	logger := log.New(os.Stdout, "go-openai: ", log.LstdFlags)

	client := &OpenAIClient{
		taskQueue:    make(chan func() error),
		RateLimit:    0,
		logger:       logger,
		APIKey:       os.Getenv("OPENAI_API_KEY"),
		Organization: os.Getenv("OPENAI_ORG"),
	}

	if client.APIKey == "" {
		// TODO: Implement Debug, Warning, Error, Fatal, Panic
		client.logger.Printf("OPENAI_API_KEY is not set.")
		return nil
	}

	return client
}

func (o *OpenAIClient) RequestChatCompletion(body ChatCompletionRequestBody) (*ChatCompletionResponse, error) {

	byts, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(byts)

	// TODO: need to determine url and model eligiblity
	req, err := http.NewRequest("POST", CHAT_COMPLETION_URL, reqBody)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if o.Organization != "" {
		req.Header.Add("OpenAI-Organization", o.Organization)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+o.APIKey)

	res, err := o.Do(req)
	if err != nil {
		return nil, err
	}

	// Need to do some sort of error checking here because it does not throw errors on non 200 statuses.
	if res.StatusCode != http.StatusOK {
		fmt.Printf("\nThe response error is: %+v", res)
		return nil, errors.New("Invalid request.")
	}

	var resp ChatCompletionResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
