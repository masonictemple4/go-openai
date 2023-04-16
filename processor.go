package main

import (
	"os"
	"strconv"
	"strings"
)

// Might want to think about how to return an error here.
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

func ProcessCommand(cmd, curMode string) (string, error) {
	colonIndex := strings.Index(cmd, ":")

	var mode string
	if colonIndex == 0 {
		mode = strings.Split(cmd, ":")[1]
	}

	switch mode {
	case "settings":
		// TODO: Display settings menu and set some sort of global mode.
		println("please choose a settings mode")
		return mode, nil
	case "exit":
		os.Exit(-1)
	default:
		resp := "default"
		return resp, nil
	}
	return "", nil
}
