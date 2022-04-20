package concurrencycoding

import (
	"fmt"
	"sync"
	"testing"
)

func TestPipeline(t *testing.T) {
	for num := range sq(sq(gen(2, 3, 4))) {
		fmt.Println(num)
	}
	for num := range sq(sq(gen2(2, 3, 4))) {
		fmt.Println(num)
	}
}

func TestMergePipeline(t *testing.T) {
	ch := gen(2, 3, 4)
	sq1 := sq(ch)
	sq2 := sq(ch)
	out := merge(sq1, sq2)
	for num := range out {
		fmt.Println(num)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

// When the number of values to be sent is known at channel creation time, a buffer can simplify the code.
func gen2(nums ...int) chan int {
	out := make(chan int, len(nums))
	for _, num := range nums {
		out <- num
	}
	close(out)
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func merge(ch ...chan int) <-chan int {
	wg := sync.WaitGroup{}
	out := make(chan int)
	output := func(c <-chan int) {
		for num := range c {
			out <- num
		}
		wg.Done()
	}
	wg.Add(len(ch))
	for _, in := range ch {
		go output(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
