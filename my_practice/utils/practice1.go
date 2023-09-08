package utils

import (
	"fmt"
	"math"
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


// The expression T(v) converts the value v to the type T.

// Unlike in C, in Go assignment between items of different type
// requires an explicit conversion.
// Try removing the float64 or uint conversions in the example and see what happens.

func Practice1_2() {
	var x, y = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}


func Practice1_3() {
	v := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}

const Pi = 3.14

func Practice1_4() {
	const World = "세계"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	// Big is the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or binary number 10 = 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func Practice1_5() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat((Big)))
}