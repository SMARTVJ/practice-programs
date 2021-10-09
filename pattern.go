package main

import "fmt"

func apattern(row int) {

	for i := 1; i <= row; i++ {
		for j := 1; j <= row-i; j++ {
			fmt.Print(".")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}

func dpattern(row int) {
	for i := 1; i <= row; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(".")
		}

		for k := 0; k <= row-i; k++ {

			fmt.Print("* ")
		}
		fmt.Println()
	}

}

func main() {

	row := 5

	apattern(row)
	dpattern(row)

}
