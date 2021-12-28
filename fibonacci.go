//fibonacci series (n = (n-1)+(n-2)) (0,1,1,2,3,5,8,13...)

package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 || n == 0 {
		return n
	} else {

		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {

	for j := 0; j <= 10; j++ {

		fmt.Print(fibonacci(j), ",")

	}

}
