package main

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	// TODO sum a
	// TODO send result on res
	tot := 0
	for i := 0; i < len(a); i++ {
		tot += a[i]
	}
	// fmt.Println(tot)
	res <- tot
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	x, y := <-ch, <-ch
	// fmt.Println(x + y)
	// TODO Get the subtotals from the channel and return their sum
	return x + y
}
