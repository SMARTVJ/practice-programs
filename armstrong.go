// find ARmstrog number (a number that is equal to some of cube of its own digits)
package main

import "fmt"

func armstrong(num int) {
	var remainder, temp, sum int

	temp = num
	for {
		remainder = temp % 10
		sum += remainder * remainder * remainder
		temp /= 10
		if temp == 0 {
			break
		}
	}

	if sum == num {
		fmt.Println("the number is armstrong")
	} else {

		fmt.Println("not armstrong")
	}
}

func main() {

	armstrong(153) // 153=1*1*1* +5*5*5 + 3*3*3
}
