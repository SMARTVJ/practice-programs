// sort a character of "I am Developer."

package main

import "fmt"

func sort(s string) string {
	v := []rune(s)
	n := len(v) - 1
	swap := true
	for swap {
		swap = false
		for i := 0; i < n; i++ {

			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				swap = true
			}
		}
	}
	return string(v)
}

func main() {
	v := "I am Developer."
	out := sort(v)
	fmt.Println(out)
}
