package main

import "fmt"

func stringsort(v []rune) string {

	n := len(v)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				swapped = true
			}
		}

	}

	return string(v)
}

func main() {

	v := []rune("i am vijay")
	out := stringsort(v)

	fmt.Println(out)
}
