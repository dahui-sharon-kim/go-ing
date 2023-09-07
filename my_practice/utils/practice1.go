package utils

import (
	"fmt"
	"math/cmplx"
)

// // A var declaration can include initializers, one per variable.
// // If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
// var i, j int = 1, 2 // var i, j = 1, 2 가능

// func main() {
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, c, python, java)
// }

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func Practice1_1() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}