package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibs := []int{0, 1}
	count := 0
	return func() int {
		if count == 0 {
			count++
			return 0
		} else if count == 1 {
			count++
			return 1
		}
		fibs = append(fibs, fibs[len(fibs) - 2] + fibs[len(fibs) - 1])
		count++
		return fibs[len(fibs)-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
