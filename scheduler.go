package main

import (
	"fmt"
	"sync"
)

var cond = sync.NewCond(&sync.Mutex{})
var parallelism int

func scheduler(wg *sync.WaitGroup, done chan TaskResult) {
	updateParallelism(min)

	for result := range done {
		incStats(result)

		var p = optimalParallelism()
		if p < 0 {
			p = max
		} else if p < min {
			p = min
		} else if p > max {
			p = max
		}

		if p-parallelism > step {
			p = parallelism + step
		} else if parallelism-p > step {
			p = parallelism - step
		}

		updateParallelism(p)

		wg.Done()
	}
}

func updateParallelism(p int) {
	if parallelism != p {
		parallelism = p
		fmt.Printf("Updating parallelism to %d\n", p)
		cond.Broadcast()
	}
}
