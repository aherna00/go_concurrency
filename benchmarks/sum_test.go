package benchmarks

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

const arraySize = 10000000 //1 mil size array

var array = make([]int, arraySize)

func init() {

	//runtime.GOMAXPROCS(1)

	seed := rand.NewSource(time.Now().UnixNano())

	r := rand.New(seed)

	for i := 0; i < arraySize; i++ {
		array[i] = r.Intn(1000)
	}
}

func BenchmarkSumOfSums(b *testing.B) {

	b.Run("BenchmarkSingleThread", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

			sum := 0
			for _, v := range array {
				sum += v * v
			}
		}
	})

	b.Run("BenchmarkConcurrent", func(b *testing.B) {
		const numGoRoutines = 8

		for i := 0; i < b.N; i++ {

			chunkSize := len(array) / numGoRoutines

			var wg sync.WaitGroup
			sumChan := make(chan int, numGoRoutines)

			for n := 0; n < numGoRoutines; n++ {

				wg.Add(1)

				go func(start int) {
					defer wg.Done()
					sum := 0
					end := start + chunkSize
					if end > len(array) {
						end = len(array)
					}

					for _, v := range array[start:end] {
						sum += v * v
					}

					sumChan <- sum

				}(n * chunkSize)
			}

			go func() {
				wg.Wait()
				close(sumChan)
			}()

			totalSum := 0

			for sum := range sumChan {
				totalSum += sum
			}

		}
	})

}
