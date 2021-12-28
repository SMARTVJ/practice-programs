// remove dublicate number in array

package main

import "fmt"

func removeDublicate(arr []int) {
	out := []int{}
	key := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if key[arr[i]] == false {
			key[arr[i]] = true
			out = append(out, arr[i])
		}
	}
	fmt.Println(out)
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 2, 3, 7, 8, 9, 6}
	removeDublicate(arr)

}
