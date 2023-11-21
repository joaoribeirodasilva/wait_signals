package main

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/joaoribeirodasilva/wait_signals"
)

// Example
// SleepWait(sleep_time time.Duration, sigs ...os.Signal) os.Signal.
func main() {

	fmt.Println("wait_signals.SleepWait example")
	fmt.Println("wait 5 seconds for a sleep timeout or press CRTL+C to exit by a signal.")
	// the thread will block until a timeout set by the *sleep_time*
	// parameter or until it gets a syscall.SIGINT or a syscall.SIGTERM
	// signal (Ex: CTRL+C).
	sig := wait_signals.SleepWait(
		time.Duration(5000)*time.Millisecond, // sleep for 5 seconds.
		syscall.SIGINT,                       // wait for syscall.SIGINT.
		syscall.SIGTERM,                      // or for syscall.SIGTERM.
	)

	// if *sig* is null the we unblock due to the duration timeout.
	// if *sig* is not nil the it's a pointer to the signal received.
	if sig == nil {
		fmt.Printf("we exited due to a timeout\n")
	} else if *sig == syscall.SIGINT {
		fmt.Printf("we exited due to syscall.SIGINT signal\n")
	} else if *sig == syscall.SIGTERM {
		fmt.Printf("we exited due to syscall.SIGTERM signal\n")
	}

	os.Exit(0)
}
