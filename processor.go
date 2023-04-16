package main

import (
	"os"
	"strings"
)

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
