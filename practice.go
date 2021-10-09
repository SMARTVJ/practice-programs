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

func sorting(str []rune) string {
	n := len(str)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {

			if str[i] > str[i+1] {
				str[i], str[i+1] = str[i+1], str[i]
				swapped = true
			}
		}
	}
	return string(str)
}

func prime(p int) {

	isprime := true
	for i := 2; i <= p/2; i++ {

		if p%i == 0 {
			fmt.Println("this is not prime")
			isprime = false
			break
		}
	}
	if isprime == true {

		fmt.Println("prime")
	}

}

func prime2(num1, num2 int) {

	for num1 < num2 {
		isprime := true
		for i := 2; i <= num1/2; i++ {

			if num1%i == 0 {
				isprime = false
				break
			}
		}
		if isprime == true {
			fmt.Printf("%d,", num1)

		}
		num1++
	}

}

func reverse(r string) {

	for i := len(r) - 1; i >= 0; i-- {

		fmt.Printf("%c", r[i])
	}
}
func main() {

	v := []int{6, 8, 9, 3, 2}
	out := sort(v)
	fmt.Println(out)
	str := []rune("i am vijay")

	out2 := sorting(str)
	fmt.Println(out2)

	p := 48
	prime(p)

	prime2(1, 100)
	r := "i am vijay"
	reverse(r)
}
