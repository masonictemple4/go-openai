package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	conf "github.com/masonictemple4/go-openai/internal/config"
	"github.com/masonictemple4/go-openai/internal/openai"
)

var MODE string = "default"

func main() {

	apiMode := flag.String("mode", "chat", "Chat, Completion, Image, Audio, File")
	config := flag.String("l", "/etc/env/openai.env", "Config path.")
	model := flag.String("m", openai.MODEL_GPT4, "language model to use.")
	role := flag.String("r", openai.OpenAIRoleUser, "The role of the completion.")
	maxTokens := flag.Int("t", 2048, "Max tokens to use.")
	numChoices := flag.Int("c", 1, "How many choices/variations would you like.")
	user, _ := user.Current()
	username := flag.String("u", user.Username, "If you have a preferred username to be called by otherwise this defaults to your system username.")

	flag.Parse()

	// Load env
	conf.LoadConfig(*config)

	client := openai.New(nil)

	fmt.Printf("\nThe model is: %s\n", *model)

	reader := bufio.NewReader(os.Stdin)

	// TODO: Should probably move this over to the menu
	// file.
	fmt.Println("Welcome to ChatGPT for the terminal.")
	fmt.Println("Brought to you by: masonictemple4")
	fmt.Println("------------------------------------")

	done := make(chan bool)

	for {
		fmt.Printf("\n%s: ", *username)
		text, _ := reader.ReadString('\n')

		// convert clrf to lf
		text = strings.Replace(text, "\n", "", -1)
		if len(text) > *maxTokens {
			fmt.Printf("Your max limit is %d, you entered %d. Please try again.", *maxTokens, len(text))
		}

		if openai.ContainsCommand(text) {
			cmd, err := openai.ToCommand(text)
			if err != nil {
				log.Fatal(err)
			}
			err = openai.ProcessCommand(*cmd)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}

		if *apiMode == "image" {
			println()
			go openai.DisplayLoading("ChatGPT: Processing your image request", done)
			prompt, size := openai.ProcessImagePrompt(text)
			newReq := openai.ImageGenerationRequestBody{
				Size:           fmt.Sprintf("%dx%d", size, size),
				Prompt:         prompt,
				ResponseFormat: "url",
				N:              int64(*numChoices),
			}

			response, err := client.RequestImageGeneration(newReq)
			if err != nil {
				log.Fatal(err)
			}

			done <- true
			println()

			fmt.Printf("\nChat%s: %s\n", strings.ToUpper(*model), response.CleanText())
			continue
		}

		println()
		go openai.DisplayLoading("ChatGPT: Processing your request", done)
		// TODO: We're going to need to detect what model to use here
		// based on the completion model selected.
		newReq := openai.ChatCompletionRequestBody{
			Model:          *model,
			MaxTokens:      int64(*maxTokens),
			Messages:       []openai.ChatCompletionRequestMessage{{Role: *role, Content: text}},
			N:              int64(*numChoices),
			ResponseFormat: openai.ChatCompletionRequestResponseFormat{Type: openai.ResponseFormatTypeText},
		}

		// TODO: We're going to want to also make this an interface instead so it can take either type.
		response, err := client.RequestChatCompletion(newReq)
		if err != nil {
			log.Fatal(err)
		}

		done <- true
		println()

		// Nice for debugging later.
		// data, _ := json.MarshalIndent(response, "", "    ")
		// fmt.Printf("\nResponseJSON: %s\n", string(data))

		fmt.Printf("\nChat%s: %s\n", strings.ToUpper(*model), response.CleanText())

	}

}
