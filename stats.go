package main

import (
	"math"
	"time"
)

var stats struct {
	start    time.Time
	deadline time.Time

	count float64
	done  float64
	mean  float64
}

func initStats(count int) {
	stats.count = float64(count)
	stats.start = time.Now()
	stats.deadline = stats.start.Add(duration)
}

func incStats(res TaskResult) {
	stats.done += 1.0
	seconds := res.duration.Seconds()
	delta := seconds - stats.mean
	stats.mean += delta / stats.done
}

func estimateDuration() float64 {
	return stats.mean
}

func optimalParallelism() int {
	var mean = estimateDuration()
	var optimum = (stats.count - stats.done) * mean / time.Until(stats.deadline).Seconds()

	return int(math.Ceil(optimum))
}
