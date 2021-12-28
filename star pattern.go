/*
	*
       * *
      * * *
     * * * *
    * * * * *

*/

package main

import "fmt"

func main() {
	num := 5
	for i := 1; i <= num; i++ {
		for j := 1; j <= num-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i-1; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
