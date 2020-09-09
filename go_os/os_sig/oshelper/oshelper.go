package oshelper

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RegisterSignal(signals ...os.Signal) *chan os.Signal {

	doneChannel := make(chan os.Signal, 1)
	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGABRT)
	go func() {
		fmt.Println("Waiting")
		sig := <-osChan
		doneChannel <- sig
	}()

	return &doneChannel
}
