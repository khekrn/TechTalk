package falsesharing

import (
	"sync"
	"testing"

	"golang.org/x/sys/cpu"
)

type PlaceHolder struct {
	n int
}

const iter = 1000

var result int

func BenchmarkIteration(b *testing.B) {
	structA := PlaceHolder{} // Initialization
	structB := PlaceHolder{} // Initialization
	wg := sync.WaitGroup{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() { // Spin up first goroutine
			for j := 0; j < iter; j++ {
				structA.n += j
			}
			wg.Done()
		}()
		go func() { // Spin up second goroutine
			for j := 0; j < iter; j++ {
				structB.n += j
			}
			wg.Done()
		}()
		wg.Wait()                      // Wait
		result = structA.n + structB.n // Aggregate
	}
}

func BenchmarkIterationCommunication(b *testing.B) {
	ch := make(chan int, 2)
	for i := 0; i < b.N; i++ {
		go func() { // Spin up first goroutine
			i := 0 // Local state
			for j := 0; j < iter; j++ {
				i += j
			}
			ch <- i
		}()
		go func() { // Spin up second goroutine
			i := 0 // Local state
			for j := 0; j < iter; j++ {
				i += j
			}
			ch <- i
		}()
		result = <-ch + <-ch // Wait and aggregate
	}
}

type PaddedStruct struct {
	_ cpu.CacheLinePad
	n int
	_ cpu.CacheLinePad
}

func BenchmarkIterationWithPadding(b *testing.B) {
	structA := PaddedStruct{} // Initialization
	structB := PaddedStruct{} // Initialization
	wg := sync.WaitGroup{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() { // Spin up first goroutine
			for j := 0; j < iter; j++ {
				structA.n += j
			}
			wg.Done()
		}()
		go func() { // Spin up second goroutine
			for j := 0; j < iter; j++ {
				structB.n += j
			}
			wg.Done()
		}()
		wg.Wait() // Wait
	}
}
