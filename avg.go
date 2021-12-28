//average= sum/n

package main

import "fmt"

func main() {
	num := []float32{5, 6, 9, 5, 5}

	var sum float32
	for i := 0; i <= len(num)-1; i++ {
		sum += num[i]

	}
	n := float32(len(num))
	avg := sum / n
	fmt.Println(avg)
}
