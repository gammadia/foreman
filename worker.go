package main

import (
	"os"
	"os/exec"
	"time"
)

type TaskResult struct {
	duration time.Duration
}

func worker(id int, queue chan []string, done chan TaskResult) {
	for {
		cond.L.Lock()
		for parallelism < id {
			cond.Wait()
		}
		cond.L.Unlock()

		var task = <-queue

		var start = time.Now()
		var result = TaskResult{}

		cmd := exec.Command(task[0], task[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

		result.duration = time.Since(start)
		done <- result
	}
}
