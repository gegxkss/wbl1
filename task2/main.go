package main

import (
	"fmt"
	"sync"
)

func main() {
	num := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for i, number := range num {
		wg.Add(1)
		go func(idx int, n int) {
			defer wg.Done()
			square := n * n
			fmt.Printf("Горутина %d: %d^2 = %d\n", idx, n, square)
		}(i, number)
	}

	wg.Wait()
	fmt.Println("Программа завершена")
}
