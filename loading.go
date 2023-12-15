package main

import (
	"fmt"
	"strings"
	"time"
)

func DisplayLoading(msg string, ch chan bool) {
	ticker := time.NewTicker(500 * time.Millisecond)
	i := 0
	for {
		select {
		case <-ch:
			ticker.Stop()
			return
		case <-ticker.C:
			i++
			fmt.Printf("%s %s\r", msg, strings.Repeat(".", i))
		}
	}
}
