package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"

	"github.com/joho/godotenv"
)

type OpenAIClient struct {
	Organization string
	APIKey       string
	RateLimit    int64
}

func (o OpenAIClient) RequestCompletion(body ChatCompletionRequestBody) (*ChatCompletionResponse, error) {
	client := &http.Client{}

	byts, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(byts)

	// TODO: need to determine url and model eligiblity
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", reqBody)
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

var MODE string = "default"

func main() {

	config := flag.String("l", "/etc/env/openai.env", "Config path.")
	model := flag.String("m", "gpt-3.5-turbo", "language model to use.")
	role := flag.String("r", ChatCompletionMessageRoleMap["user"], "The role of the completion.")
	maxTokens := flag.Int("t", 2048, "Max tokens to use.")
	numChoices := flag.Int("c", 1, "How many choices/variations would you like.")
	user, _ := user.Current()
	username := flag.String("u", user.Username, "If you have a preferred username to be called by otherwise this defaults to your system username.")
	//
	flag.Parse()

	// Load env
	LoadConfig(*config)

	newAiClient := OpenAIClient{
		Organization: os.Getenv("OPENAI_ORG"),
		APIKey:       os.Getenv("OPENAI_API_KEY"),
	}

	fmt.Printf("\nThe model is: %s\n", *model)

	reader := bufio.NewReader(os.Stdin)

	// TODO: Should probably move this over to the menu
	// file.
	fmt.Println("Welcome to ChatGPT for the terminal.")
	fmt.Println("Brought to you by: masonictemple4")
	fmt.Println("------------------------------------")

	for {
		fmt.Printf("\n%s: ", *username)
		text, _ := reader.ReadString('\n')

		// convert clrf to lf
		text = strings.Replace(text, "\n", "", -1)
		if len(text) > *maxTokens {
			fmt.Printf("Your max limit is %d, you entered %d. Please try again.", *maxTokens, len(text))
		}

		MODE, err := ProcessCommand(text, MODE)
		if err != nil {
			log.Fatal(err)
		}

		if MODE == "default" {
			// TODO: We're going to need to detect what model to use here
			// based on the completion model selected.
			newReq := ChatCompletionRequestBody{
				Model:     *model,
				MaxTokens: int64(*maxTokens),
				Messages:  []ChatCompletionRequestMessage{{Role: *role, Content: text}},
				N:         int64(*numChoices),
			}

			// TODO: We're going to want to also make this an interface instead so it can take either type.
			response, err := newAiClient.RequestCompletion(newReq)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("\nChatGPT: %s\n", FormatChatResponseText(response))
		}

	}

}

// TODO: Needs to be an interface function for the Response models.
func FormatChatResponseText(resp *ChatCompletionResponse) string {
	text := resp.Choices[0].Message.Content
	text = strings.TrimLeft(text, "?\n")
	return text
}

func FormatResponseText(resp *CompletionResponse) string {
	text := resp.Choices[0].Text
	text = strings.TrimLeft(text, "?\n")
	return text
}

func LoadConfig(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
