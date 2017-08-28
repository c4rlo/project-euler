package main

import "fmt"

func main() {
	var fifthPowers [10]int
	for i := 1; i <= 9; i++ {
		fifthPowers[i] = i * i * i * i * i
	}

	total := 0
	for n := 2; n <= 999999; n++ {
		s := 0
		for m := n; m > 0; {
			s += fifthPowers[m%10]
			m /= 10
		}
		if n == s {
			fmt.Println(n)
			total += n
		}
	}
	fmt.Println("Total:", total)
}
