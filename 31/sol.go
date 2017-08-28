package main

import "fmt"

var Coins = [...]int{200, 100, 50, 20, 10, 5, 2, 1}

func numWays(amount int, coins []int) int {
	n := 0
	for i, c := range coins {
		if c == amount {
			n++
		} else if c < amount {
			n += numWays(amount-c, coins[i:])
		}
	}
	return n
}

func main() {
	fmt.Println("Result:", numWays(200, Coins[:]))
}
