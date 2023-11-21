# Wait Signal

This library has a BSD 3-Clause License.

Use it has you will.

## Disclaimer

See: The [LICENCE](https://github.com/joaoribeirodasilva/wait_signals/blob/main/LICENSE) information.

## Overview

The Go module **wait_signal** implements two widely used blocks of code for waiting for [os.Signal](https://pkg.go.dev/os#Signal)s from the operating system.

Both functions wait until one of the desired [os.Signal](https://pkg.go.dev/os#Signal)s to be fired by the operating system but another also waits for a certain timeout.

* The **Wait** function just blocks the thread execution until it get one of the desired signals.
* The **SleepWait** function blocks the thread execution until it get one of the desired signals or a timeout occurs.

## Add to your project

```bash
go get github.com/joaoribeirodasilva/wait_signal
```

## Usage

Bellow you will find the usage exemples.

### SleepWait() Example

```go
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
        time.Duration(5000) * time.Millisecond, // sleep for 5 seconds.
        syscall.SIGINT, // wait for syscall.SIGINT.
        syscall.SIGTERM, // or for syscall.SIGTERM.
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
```

### Wait() Example

```go
package main

// Example
// Wait(sigs ...os.Signal) *os.Signal.
func main() {

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
```
