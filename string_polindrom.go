package main

import "fmt"

func reverse(str string) string {
	sum := ""
	for i := len(str) - 1; i >= 0; i-- {

		fmt.Printf("%c", str[i])

		sum += string(str[i])
	}
	fmt.Println()
	return sum
}

func main() {

	str := "viv"

	s := reverse(str)
	if str == s {

		fmt.Println(str, "is polindrom")
	} else {
		fmt.Println(str, "is not polindrom")

	}

}
