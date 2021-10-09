package main

import (
	"fmt"
)

func main() {
	num := 17

	isprime := true
	for i := 2; i <= num/2; i++ {

		if num%i == 0 {
			fmt.Printf("the %d is not prime", num)
			isprime = false
			break
		}
	}
	if isprime == true {
		fmt.Printf("the %d is prime number", num)
	}

}
