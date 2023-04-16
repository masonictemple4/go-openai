package main

import (
	"errors"
	"strconv"
	"strings"
)

// create a data type to represent the following parts of this string :set mode image
type Command struct {
	// This will be set, get, help
	Verb string
	// The setting to configure
	// 	apiMode, config, model, role, maxTokens, numChoices, username
	Setting string
	// Currently only available for env settings.
	SubSetting string
	// The subsetting to configure
	// Example: apikey, orgid
	Value string
}

// ContainsCommand will check the input text to ensure that it contains a command.
func ContainsCommand(text string) bool {
	colonIndex := strings.Index(text, ":")

	return colonIndex == 0
}

// ToCommand will return a pointer to a new command from the text given to it.
// Or an error if it's in the incorrect format.
// Note: this can really be called mainly after ContainsCommand returns true.
func ToCommand(text string) (*Command, error) {
	if strings.Contains(text, ":") == false {
		return nil, errors.New("Invalid command")
	}
	parts := strings.Split(strings.Split(text, ":")[1], " ")

	if len(parts) > 4 {
		return nil, errors.New("Invalid command")
	}

	if len(parts) == 4 {
		return &Command{
			Verb:       parts[0],
			Setting:    parts[1],
			SubSetting: parts[2],
			Value:      parts[3],
		}, nil
	}
	return &Command{
		Verb:    parts[0],
		Setting: parts[1],
		Value:   parts[2],
	}, nil
}

func ProcessCommand(cmd Command) error {
	return nil
}

// Might want to think about how to return an error here.
// The default is going to be 256x256 if the size fails to parse.
func ProcessImagePrompt(text string) (string, int64) {
	colonIndex := strings.Index(text, ":")

	// Also check length here when refactoring.
	sizeStr := strings.Split(strings.Split(text, ":")[1], " ")[1]
	size, _ := strconv.ParseInt(sizeStr, 10, 64)

	// Set a default if it fails.
	if size == 0 {
		size = 256
	}

	return strings.TrimSpace(text[:colonIndex]), size
}
