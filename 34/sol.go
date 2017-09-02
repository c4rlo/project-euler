package main

import "fmt"

func main() {
	var factorial [10]int
	factorial[0] = 1
	k := 1
	for i := 1; i <= 9; i++ {
		k *= i
		factorial[i] = k
	}

	total := 0

	for n := 3; n < 10000000; n++ {
		f := 0
		for k := n; k > 0; {
			f += factorial[k%10]
			k /= 10
		}
		if n == f {
			fmt.Println(n)
			total += n
		}
	}

	fmt.Printf("\nTotal: %v\n", total)
}
