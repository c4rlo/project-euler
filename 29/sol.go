package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var PrimesBelowTen = [...]int{2, 3, 5, 7}

type PrimeFactor struct {
	prime    int
	exponent int
}

type PrimeFac []PrimeFactor

func (pf *PrimeFac) addNextPrimeFactor(p int) {
	// 'p' must be >= the largest existing prime factor in 'pf'
	if *pf != nil && (*pf)[len(*pf)-1].prime == p {
		(*pf)[len(*pf)-1].exponent++
	} else {
		*pf = append(*pf, PrimeFactor{p, 1})
	}
}

func (pf *PrimeFac) Pow(result *PrimeFac, e int) {
	*result = make(PrimeFac, len(*pf))
	copy(*result, *pf)
	for i := range *result {
		(*result)[i].exponent *= e
	}
}

func (pf *PrimeFac) Serz() string {
	var buf bytes.Buffer
	for _, pf := range *pf {
		if err := binary.Write(&buf, binary.LittleEndian, uint8(pf.prime)); err != nil {
			panic(err)
		}
		if err := binary.Write(&buf, binary.LittleEndian, uint16(pf.exponent)); err != nil {
			panic(err)
		}
	}
	// fmt.Printf("  pf = %v; len(serz) = %v\n", pf, len(buf.Bytes()))
	return string(buf.Bytes())
}

func primeFac(n int) PrimeFac {
	if n > 100 {
		panic("n > 100 not supported")
	}
	var result PrimeFac = nil
	for _, p := range PrimesBelowTen {
		for n%p == 0 {
			result.addNextPrimeFactor(p)
			n /= p
		}
	}
	if n > 1 {
		result.addNextPrimeFactor(n)
	}
	return result
}

func main() {
	values := make(map[string]bool)

	for a := 2; a <= 100; a++ {
		aFac := primeFac(a)
		var value PrimeFac
		for b := 2; b <= 100; b++ {
			aFac.Pow(&value, b)
			// fmt.Printf("%v(%v)^%v = %v\n", a, aFac, b, value)
			values[value.Serz()] = true
		}
	}
	fmt.Printf("map size = %v\n", len(values))
}
