// Copyright 2023 João Ribeiro da Silva. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wait_signals

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

// signals to send
var sendSignals = []os.Signal{syscall.SIGALRM, syscall.SIGCONT}

// base waiting time
const WAIT_TIME = 100

// TestWait tests the signals Wait function
func TestWait(t *testing.T) {

	// get our pid
	pid := os.Getpid()

	// finds the process
	process, err := os.FindProcess(pid)
	if err != nil {
		t.Fatalf("error fetching process for pid %d, %s", pid, err.Error())
	}

	// flags the main thread
	isChild := false

	// new thread to send the signal to us
	go func() {

		// sets this thread to be a child one, not the main thread
		isChild = true

		// waits a bit so the main thread can wait for our signal
		time.Sleep(WAIT_TIME * time.Millisecond)

		// sends the signal to the main thread that is waiting for it
		err = process.Signal(sendSignals[0])

	}()

	// checks if the thread failed to send the signal
	// the err will be nil on the main thread
	if err != nil {
		t.Fatalf("error sending %d signal, %s", sendSignals[0], err.Error())
	}

	// check if this thread is not the go func one
	if !isChild {

		// main thread waits to receive the signal send
		sig := Wait(sendSignals...)

		// if the returned signal is not the same we sent
		// report an error
		if sig != sendSignals[0] {
			t.Fatalf("received signal %d should have received %s", sig, sendSignals[0])
		}

		signal.Reset(sendSignals[0])
	}

}

func TestSleepWaitTimeout(t *testing.T) {

	// stores the start time
	start := time.Now()

	// sleeps for a certain time
	sig := SleepWait(WAIT_TIME*time.Millisecond, sendSignals...)
	if sig != nil {
		t.Fatalf("received %d should receive %d", *sig, sendSignals[0])
	}

	// stores the end time
	end := time.Now()

	// checks the total time slept
	total := int64(end.Sub(start) / time.Millisecond)

	// if the time slept is smaller that the time it should sleep
	// return an error
	if total < WAIT_TIME {
		t.Fatalf("received %d should receive %d", *sig, sendSignals[0])
	}
}

func TestSleepWaitSignal(t *testing.T) {

	// get our pid
	pid := os.Getpid()

	// finds the process
	process, err := os.FindProcess(pid)
	if err != nil {
		t.Fatalf("error fetching process for pid %d, %s", pid, err.Error())
	}

	// flags the main thread
	isChild := false

	// new thread to send the signal to us
	go func() {

		// sets this thread to be a child one, not the main thread
		isChild = true

		// waits a bit so the main thread can wait for our signal
		time.Sleep(WAIT_TIME * time.Millisecond / 2)

		// sends the signal to the main thread that is waiting for it
		err = process.Signal(sendSignals[1])

	}()

	// checks if the thread failed to send the signal
	// the err will be nil on the main thread
	if err != nil {
		t.Fatalf("error sending %d signal, %s", sendSignals[1], err.Error())
	}

	// check if this thread is not the go func one
	if !isChild {

		// main thread waits to receive the signal send
		sig := SleepWait(WAIT_TIME*time.Millisecond, sendSignals...)

		// if sig is nill it means the a sleep timeout occurred and
		// not a signal exit as we expected
		if sig == nil {
			t.Fatalf("no signal received before wake from sleep")
		}

		// if the returned signal is not the same we sent
		// report an error
		if *sig != sendSignals[1] {
			t.Fatalf("received %d should receive %d", *sig, sendSignals[1])
		}

		signal.Reset(sendSignals[1])
	}

}
