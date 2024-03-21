package main

import "fmt"

var y int

func main() {
	//var for zero value
	fmt.Println(y)

	//short declaration operator
	z := 42
	fmt.Println(z)

	//multiple initializations
	a, b := 43, "Happiness"
	fmt.Println(a, b)

	// var when specificity is required
	var c float32 = 42.42
	fmt.Printf("%v is of this type %T\n", c, c)

	//blank identifier
	e, f, _ := 45, 46, 47
	fmt.Println(e, f)

}
