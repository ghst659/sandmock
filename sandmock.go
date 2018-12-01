package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type task struct {
	name string
	latency time.Duration
	duration time.Duration
}

func delay(duration time.Duration) {
	time.Sleep(duration * time.Second)
}

func (t task) run(wg *sync.WaitGroup, c io.Writer) {
	defer wg.Done()
	delay(t.latency)
	fmt.Fprintf(c, "Begin %s\n", t.name)
	delay(t.duration)
	fmt.Fprintf(c, "Finished %s\n", t.name)
}

func main() {
	tasks := []task{
		{
			name: "A:2-5",
			latency: 2,
			duration: 5,
		},
		{
			name: "B:1-13",
			latency: 1,
			duration: 13,
		},
		{
			name: "C:0-7",
			latency: 0,
			duration: 7,
		},
		{
			name: "D:3-3",
			latency: 3,
			duration: 3,
		},
		{
			name: "E:6-1",
			latency: 6,
			duration: 1,
		},
	}
	var wg sync.WaitGroup
	for i, t := range tasks {
		wg.Add(1)
		var c io.Writer
		if i % 2 != 0 {
			c = os.Stdout
		} else {
			c = os.Stderr
		}
		go t.run(&wg, c)
	}
	wg.Wait()
}
