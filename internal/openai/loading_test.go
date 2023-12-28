package openai

import (
	"testing"
	"time"
)

func TestLoading(t *testing.T) {
	ch := make(chan bool)
	go DisplayLoading("test", ch)
	time.Sleep(10 * time.Second)
	ch <- true
}
