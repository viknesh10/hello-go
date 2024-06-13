package main

import "fmt"

func variables() {
	// Different ways of declaring a variable
	var i int
	i = 42
	var j int = 27 // Declaring this way gives you better control
	k := 99        // Doesnt let you declare something like float32 for example. 99.1 will give you just float64
	// Variable grouping
	// Helpful to group domain level variables together. Not a feature or syntax but just a way to organise code better
	// var (
	// 	name       string = ""
	// 	schoolName string = ""
	// 	age        int    = 7
	// )
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
