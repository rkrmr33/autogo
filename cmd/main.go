package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hooks "github.com/robotn/gohook"
)

func main() {
	interval := flag.Int("i", 1000, "The click interval in milliseconds (default 1000 ms)")
	startChan := make(chan interface{})
	pause := true
	flag.Parse()
	fmt.Println("Press [k] to start, send any signal to stop...")

	go func() {
		evChan := hooks.Start()

		for ev := range evChan {
			if ev.Rawcode == 40 && ev.Kind == hooks.KeyUp {
				if pause {
					pause = false
					startChan <- struct{}{}
				} else {
					pause = true
				}
			}
		}
	}()

	for {
		<-startChan
		fmt.Printf("Right clicking every %v ms\n", *interval)
		for !pause {
			time.Sleep(time.Duration(*interval) * time.Millisecond)
			fmt.Println("right clicking...")
			robotgo.Click("right", false)
		}
		fmt.Println("Paused")
	}
}
