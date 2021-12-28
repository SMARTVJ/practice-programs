/*
	*
       * *
      *   *
     *     *
    *********
*/

package main

import "fmt"

func main() {
	num := 5

	for i := 1; i <= num; i++ {
		for j := 1; j <= num-i; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == num {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Println()

	}

}
