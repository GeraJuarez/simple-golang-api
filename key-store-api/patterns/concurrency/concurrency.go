package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Split(source <-chan string, n int) []<-chan string {
	dests := make([]<-chan string, 0)

	for i := 0; i < n; i++ {
		ch := make(chan string)
		dests = append(dests, ch)

		go func() {
			defer close(ch)

			for val := range source {
				ch <- val
			}
		}()
	}

	return dests
}

func Funnel(sources ...<-chan string) <-chan string {
	dest := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(sources))

	for _, ch := range sources {
		go func(c <-chan string) {
			defer wg.Done()

			for n := range c {
				dest <- n
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(dest)
	}()

	return dest
}

func testFunnel() <-chan string {
	sources := make([]<-chan string, 0)

	for i := 0; i < 3; i++ {
		ch := make(chan string)
		sources = append(sources, ch)

		go func() {
			defer close(ch)

			for i := 1; i <= 5; i++ {
				ch <- fmt.Sprint(i)
				time.Sleep(time.Second)
			}
		}()
	}

	dest := Funnel(sources...)

	return dest
}

func testSplit() {
	source := make(chan string)
	dests := Split(source, 5)

	go func() {
		for i := 0; i <= 10; i++ {
			source <- fmt.Sprint(i)
		}

		close(source)
	}()

	var wg sync.WaitGroup
	wg.Add(len(dests))

	for i, ch := range dests {
		go func(i int, d <-chan string) {
			defer wg.Done()

			for val := range d {
				fmt.Printf(val)
			}
		}(i, ch)
	}

	wg.Wait()
}
