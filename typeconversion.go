package main

import (
	"fmt"
	"strconv"
)

func type_conversion() {
	var i = 42
	fmt.Printf("%v, %T\n", i, i)
	var j float32
	// Type conversion
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k string
	// String conversion in Go needs a package
	// Because string(i) would convert it to an ascii string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)
}
