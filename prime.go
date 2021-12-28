// prime number between 45 to 50

package main

import "fmt"

func prime(num1, num2 int) {
	fmt.Println("THE PRIME number :")

	if num1 < 2 || num2 < 2 {
		fmt.Println("the number must be greater than 2")
	}

	for num1 <= num2 {
		isprime := true
		for i := 2; i <= num1/2; i++ {
			if num1%i == 0 {
				isprime = false
				break
			}
		}
		if isprime == true {
			fmt.Printf("%d,", num1)
		}
		num1++

	}
}
func main() {
	prime(45, 50) // a number devided by same number and 1 its called prime number
}
