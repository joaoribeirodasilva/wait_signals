// Copyright 2023 João Ribeiro da Silva. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wait_signals

// Wait signal implements waiting functions for os.Signals
// as well as a thread sleep (like time.Sleep) that awakes
// after its sleep duration or in any of the specified
// os.Signals
import (
	"os"
	"os/signal"
	"time"
)

// Wait awaits for any of the signals in the parameter sigs
// to be received
//
// if a signal is received the function returns the signal
// that was received
func Wait(sigs ...os.Signal) os.Signal {

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, sigs...)

	sig := <-signalChannel

	return sig
}

// SleepWait sleeps the duration time in sleep_time and awaits
// for any of the signals in the parameter sigs to be received
//
// if the sleep duration finishes a nil *os.Signal is returned
// otherwise a *os.Signal containing the signal received is returned
func SleepWait(sleep_time time.Duration, sigs ...os.Signal) *os.Signal {

	sleepChannel := time.After(sleep_time)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, sigs...)

	select {
	case <-sleepChannel:
		return nil
	case sig := <-signalChannel:
		return &sig
	}
}
