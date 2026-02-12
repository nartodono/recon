package ui

import (
	"fmt"
	"time"
)

type Spinner struct {
	stop chan struct{}
	done chan struct{}
}

func NewSpinner() *Spinner {
	return &Spinner{
		stop: make(chan struct{}),
		done: make(chan struct{}),
	}
}

// Start prints spinner on the same line. Call Stop() to end it.
func (s *Spinner) Start(prefix string) {
	go func() {
		defer close(s.done)
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		for {
			select {
			case <-s.stop:
				fmt.Print("\r\033[2K")
				return
			default:
				fmt.Printf("\r\033[2K%s %s", prefix, frames[i%len(frames)])
				i++
				time.Sleep(90 * time.Millisecond)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	close(s.stop)
	<-s.done
}
