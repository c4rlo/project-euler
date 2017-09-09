package main

import "fmt"

func main() {
	numSolns := make(map[int]int)
	sqrt := make(map[int]int)
	for c := 1; c <= 500; c++ {
		cc := c * c
		sqrt[cc] = c
		for a := 1; ; a++ {
			aa := a * a
			bb := cc - aa
			if bb < aa {
				break
			}
			b := sqrt[bb]
			if b == 0 {
				continue
			}
			perim := a + b + c
			if perim > 1000 {
				break
			}
			// fmt.Printf("%v^2 + %v^2 = %v^2\n", a, b, c)
			numSolns[perim]++
		}
	}

	var maxSolnsPerim, maxSolns int
	for perim, solns := range numSolns {
		if solns > maxSolns {
			maxSolns = solns
			maxSolnsPerim = perim
		}
	}

	fmt.Printf("Max number of solutions (%v) occurs for perimeter %v.\n",
		maxSolns, maxSolnsPerim)
}
