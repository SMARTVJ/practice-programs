// remove dublicate string in array
package main

import "fmt"

func remove(str []string) {
	out := []string{}
	key := make(map[string]bool)
	for i := 0; i < len(str); i++ {
		if key[str[i]] == false {
			key[str[i]] = true
			out = append(out, str[i])
		}
	}
	fmt.Println(out)

}

func main() {

	str := []string{"vijay", "hema", "amma", "hema", "vijay"}
	remove(str)
}
