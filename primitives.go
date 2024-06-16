package main

import "fmt"

func primitives() {
	var n bool = true
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	a := 8              // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // 2^3 / 2^3 = 1
}
