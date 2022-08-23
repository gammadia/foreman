package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	flag.Parse()

	if err := validateOpts(); err != nil {
		fmt.Fprintf(os.Stderr, "Configuration error: %s\n\n", err)
		flag.Usage()
		os.Exit(1)
	}

	var tasks = prepareTasks()
	var count = len(tasks)

	initStats(count)

	var queue = make(chan []string, count)
	for _, task := range tasks {
		queue <- task
	}

	var wg sync.WaitGroup
	wg.Add(count)

	var done = make(chan TaskResult, max)

	for i := 1; i <= max; i++ {
		go worker(i, queue, done)
	}

	go scheduler(&wg, done)
	wg.Wait()
}
