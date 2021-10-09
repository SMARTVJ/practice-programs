package main

import "fmt"

func prime(num1, num2 int) {
	if num1 < 2 || num2 < 2 {
		fmt.Println("number must be greater than 0")
		return
	}

	for num1 <= num2 {
		isprime := true
		for i := 2; i <= (num1 / 2); i++ {
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
	fmt.Println()
}
func main() {

	prime(2, 100)
	prime(45, 50)

}
