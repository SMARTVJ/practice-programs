package main

import "fmt"

func sort(v []int) []int {
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

	return v
}
func main() {

	v := []int{5, 4, 9, 7}

	out := sort(v)
	fmt.Println(out)
}
