package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/joaoribeirodasilva/wait_signals"
)

// Example
// Wait(sigs ...os.Signal) *os.Signal.
func main() {

	fmt.Println("wait_signals.Wait example")
	fmt.Println("press CRTL+C to exit by a signal.")

	// the thread will block until it gets a syscall.SIGINT
	// or a syscall.SIGTERM signal (Ex: CTRL+C).
	sig := wait_signals.Wait(syscall.SIGINT, syscall.SIGTERM)

	// the sig return which signal was received.
	if sig == syscall.SIGINT {
		fmt.Printf("we received a syscall.SIGINT signal\n")
	} else {
		fmt.Printf("we received a syscall.SIGTERM signal\n")
	}
	os.Exit(0)
}
