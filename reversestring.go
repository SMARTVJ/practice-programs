package main

import "fmt"

func reverse(v string) {

	for i := len(v) - 1; i >= 0; i-- {

		fmt.Printf("%c", v[i])
	}
}

func main() {

	v := "vijay"

	reverse(v)
}
