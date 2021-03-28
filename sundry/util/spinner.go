package util

import (
	"fmt"
	"time"
)

type status struct {
	arr      []string
	stopChan chan struct{}
}

func (s *status) progress() {
	s.stopChan = make(chan struct{}, 1)

	for {
		for _, val := range s.arr {
			select {
			case <-s.stopChan:
				fmt.Printf("\r")
				return
			default:
				fmt.Printf("\r%s ... %s", "Waiting", val)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}
}

func progress(message string, display []string) chan struct{} {
	ch := make(chan struct{}, 1)

	return ch
}
