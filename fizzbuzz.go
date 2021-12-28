/*
To print "FIZZ" for multiples of 3
To print "BUZZ" for multiples of 5
To print "FIZZ BUZZ" for multiples of 3 & 5
*/
package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print(" FIZZBUZZ ")
		} else if i%3 == 0 {
			fmt.Print(" FIZZ ")
		} else if i%5 == 0 {

			fmt.Print(" BUZZ ")
		} else {
			fmt.Print(i, " ")
		}

	}

}
