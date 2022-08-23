package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var duration time.Duration
var min, max, step int
var token string

func init() {
	flag.DurationVar(&duration, "deadline", 0, "...")
	flag.IntVar(&min, "min", 1, "minimum number of parallel jobs")
	flag.IntVar(&max, "max", runtime.NumCPU(), "maximum number of parallel jobs")
	flag.IntVar(&step, "step", 1, "...")
	flag.StringVar(&token, "token", "{}", "...")
}

func validateOpts() error {
	switch true {
	case duration <= 0:
		return fmt.Errorf("-deadline missing")
	case min < 1:
		return fmt.Errorf("-min must be >= 1")
	case max < min:
		return fmt.Errorf("-max must be >= -min")
	case step < 1:
		return fmt.Errorf("-step must be >= 1")
	default:
		return nil
	}
}

func prepareTasks() [][]string {
	var tasks = [][]string{}

	var args = flag.Args()
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var cmd = make([]string, len(args))
		var line = scanner.Text()
		for i, arg := range args {
			cmd[i] = strings.ReplaceAll(arg, token, line)
		}

		tasks = append(tasks, cmd)
	}

	return tasks
}
