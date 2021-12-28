//find 1st,2nd,3rd LARGEST number 

package main

import "fmt"

func main() {

	arr := []int{69, 42, 2, 9}
	var lar1, lar2, lar3 int
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] > lar1 {
			lar2 = lar1
			lar1 = arr[i]
		} else if lar2 < arr[i] {
			fmt.Print(lar2)
			lar3 = lar2
			lar2 = arr[i]
		} else if lar3 < arr[i] {
			lar3 = arr[i]
		}
	}
	fmt.Println(lar1)
	fmt.Println(lar2)
	fmt.Println(lar3)

}
