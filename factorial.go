//factorial 3! = 3*2*1 = 6
package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {

		return n * factorial(n-1)
	}
}

func main() {
	fmt.Println(factorial(3))
}
