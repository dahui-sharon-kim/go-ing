package utils

import (
	"fmt"
	"math"
)

// Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument
// The receiver appears in its own argument list btw the func keyword and the method name.
// In this example, the Abs method has a receiver of type Vertex named v.
type Vertex struct {
	X, Y float64
}

// ** Vertex의 메소드 Abs
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// `Abs` method can be called on values of type `Vertex`
// This means that `Abs` is a method specifically designed to work with `Vertex` values.
func Practice4_1() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// Methods are functions
// Remember: a method is just a function with a receiver argument
// Here's `Abs` written as a regular function with no change in functionality

// ** 일반 함수 Abs
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Practice4_2() {
	v := Vertex{3, 4}
	fmt.Println(AbsFunc(v))
}

// You can declare a method on non-struct types, too.
// In this example we see a numeric type MyFloat with an Abs method.

type MyFloat float64

// You can only declare a method with a receiver whose type is defined
// in the same package as the method.
// You cannot declare a method with a receiver whose type is defined
// in another package (int 같은 built-in types도 안 됨!!).

func (f MyFloat) Abs4_3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Practice4_3() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs4_3())
}

// ** Pointer recievers

// You can declare methods with pointer receivers.
// This means the receiver type has the literal syntax *T for some type T.
// (Also, T cannot itself be a pointer such as *int.)

// For example, the Scale method here is defined on *Vertex.
// Methods with pointer receivers can modify the value
// to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver,
// pointer receivers are more common than value receivers.

// Try removing the * from the declaration of the Scale function
// and observe how the program's behavior changes.
// With a value receiver, the Scale method operates on a copy of the original 
// Vertex value. (This is the same behavior as for any other function argument.)
// The Scale method must have a pointer receiver to change the
// Vertex value declared in the main function.

// v가 포인터라는 뜻이 아니고,
// v가 *Vertex 타입이라는 뜻으로, *Vertex는 Vertex타입의 어떤 값에 대한 포인터임
func (v *Vertex) Scale(f float64) { // pointer receiver
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale2(f float64) { // value receiver
	v.X = v.X * f
	v.Y = v.Y * f
}

func Practice4_4() {
	v := Vertex{3, 4}
	v2 := Vertex{3, 4}
	v.Scale(10)
	// 아래에서 v.Scale2(10)을 호출할 때 v2값의 복사본을 만든 후 넘김
	// 따라서 Scale2 method에서 어떤 변화를 줘도 원본 v에는 영향이 없음
	v2.Scale2(10)
	fmt.Println(v.Abs())
	fmt.Println(v2.Abs())
}

// Pointers and functions

// Here we see the Abs and Scale methods rewritten as functions.
// Again, try removing the * from line 16.
// Can you see why the behavior changes?
// What else did you need to change for the example to compile?
// (If you're not sure, continue to the next page.)

func ScaleFunc(v *Vertex, f float64) { // argument로 pointer를 받음
	v.X *= f
	v.Y *= f
}

func ScaleFuncVal(v Vertex, f float64) { // argument로 value를 받음
	v.X *= f
	v.Y *= f
}

func Practice4_5() {
	v := Vertex{3, 4}
	// ScaleFunc(v, 10) // Complile Error
	ScaleFuncVal(v, 10)
	fmt.Println(AbsFunc(v))
	ScaleFunc(&v, 10)
	fmt.Println(AbsFunc(v))
}

// Methods and pointer indirection
// Practice4_5 보면 pointer argument를 받는 함수는 pointer 안 넣으면 컴파일 에러 발생
// 하지만 pointer recievers 받는 메소드는 value든 pointer든 뭘 넘겨도 됨
// pointer가 아니라 value를 넘겨도 pointer receiver가 자동으로 call 됨.

func Practice4_6() {
	v := Vertex{3, 4} // v는 포인터가 아님
	v.Scale(2) // {6, 8}
	ScaleFunc(&v, 10) // &v는 포인터임. {60, 80}

	p := &Vertex{4, 3} // p는 포인터
	p.Scale(3) // {12, 9}
	ScaleFunc(p, 8) // {96, 72}

	fmt.Println(v, p) // v는 포인터가 아니고 p는 포인터임
}


// Methods and pointer indirection 2


func Practice4_7() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4,3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

func Practice4_8() {
	
}
func Practice4_9() {
	
}
func Practice4_10() {
	
}
func Practice4_11() {
	
}
func Practice4_12() {
	
}
func Practice4_13() {
	
}
func Practice4_14() {
	
}
func Practice4_15() {
	
}
func Practice4_16() {
	
}
func Practice4_17() {
	
}
func Practice4_18() {
	
}
func Practice4_19() {
	
}
func Practice4_20() {
	
}
func Practice4_21() {
	
}
func Practice4_22() {
	
}
func Practice4_23() {
	
}
func Practice4_24() {
	
}
func Practice4_25() {
	
}
func Practice4_26() {
	
}