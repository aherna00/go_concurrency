package benchmarks

import (
	"sync"
	"testing"
	"time"
)

var durations []time.Duration

func init() {

	//use GOMAXPROCS to limit benchmarks to a single core
	//runtime.GOMAXPROCS(1)

	durations = []time.Duration{
		71 * time.Millisecond, 96 * time.Millisecond, 77 * time.Millisecond,
		7 * time.Millisecond, 32 * time.Millisecond, 98 * time.Millisecond,
		68 * time.Millisecond, 58 * time.Millisecond, 40 * time.Millisecond,
		61 * time.Millisecond, 25 * time.Millisecond, 96 * time.Millisecond,
		37 * time.Millisecond, 96 * time.Millisecond, 25 * time.Millisecond,
	}
}

func BenchmarkIO(b *testing.B) {

	b.Run("BenchmarkSingleCore", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			simulateIOTaskSingleThreaded(durations)
		}
	})

	b.Run("BenchmarkConcurrent", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			simulateIOTaskConcurrent(durations)
		}

	})

}

func simulateIOTask(duration time.Duration) {
	time.Sleep(duration)
}

func simulateIOTaskSingleThreaded(durations []time.Duration) {
	for _, duration := range durations {
		simulateIOTask(duration)
	}
}

func simulateIOTaskConcurrent(durations []time.Duration) {
	var wg sync.WaitGroup

	wg.Add(len(durations))

	for _, duration := range durations {

		go func(duration time.Duration) {
			defer wg.Done()
			simulateIOTask(duration)
		}(duration)
	}

	wg.Wait()
}
