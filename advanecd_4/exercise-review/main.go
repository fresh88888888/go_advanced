package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type callbackchan chan struct{}

// Traigersthe callback channel every d time.Duration units
func checkEvery(ctx context.Context, d time.Duration, cb callbackchan) {
	for {
		select {
		case <-ctx.Done():
			// ctx is canceld
			return
		case <-time.After(d):
			// wait for the duration
			if cb != nil {
				cb <- struct{}{}
			}
		}
	}
}

func printProssessList() {
	psCommand := exec.Command("ps", "a")
	resp, err := psCommand.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	out := string(resp)
	lines := strings.Split(out, "\n")

	for _, line := range lines {
		if line != "" {
			fmt.Println(line)
		}
	}
}

func main() {
	ctx := context.Background()
	printProssessList()
	callback := make(callbackchan)
	go checkEvery(ctx, time.Second*5, callback)
	go func() {
		for {
			select {
			case <-callback:
				printProssessList()
			}
		}
	}()

	for {
		time.Sleep(time.Second * 10)
	}
}
