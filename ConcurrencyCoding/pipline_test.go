package concurrencycoding

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	for num := range sq(sq(gen(2, 3, 4))) {
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
