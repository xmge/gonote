package main

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	go func() {
		s := spinner.New(spinner.CharSets[8], 100*time.Millisecond)  // Build our new spinner
		s.Start()
		time.Sleep(4 * time.Second)                                  // Run for some time to simulate work
		s.Stop()
	}()
	select {
	}
}