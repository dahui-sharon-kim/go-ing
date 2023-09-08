package utils

import (
	"fmt"
)

// Pointers

// Go has pointers. A pointer holds the memory address of a value.

// `*T`라는 타입은 T타입의 어떤 변수에 대한 포인터가 갖는 타입.
// `var p *int`는 `p`가 integer 변수의 메모리 주소를 가질 수 있다는 의미.
// 위와 같이 선언하면 `p`는 초기화되어 zero value인 `nil`의 값을 가짐

// The `&` operator generates a pointer to its operand.
// i := 42
// p = &i

// The `*` operator denotes the pointer's underlying value.
// fmt.Println(*p)
// *p = 21
// This is known as "dereferencing" or "indirecting".
// Unlike C, Go has no pointer arithmetic.

func Practice3_1() {
	i, j := 42, 2701

	// i의 pointer를 생성하려면 &i라고 쓴다
	p := &i

	// Below: READ i through the pointer p
	// pointer를 통해 실제 값을 읽으려면 *p라고 쓴다
	fmt.Println(*p)  // 42

	// Below: SET i through the pointer p
	*p = 21
	fmt.Println(i) // 21

	 // p는 j의 pointer
	p = &j

	// Below: DIVIDE j through the pointer
	*p = *p / 37   // j = j / 37
	fmt.Println(j) // 새로 set한 j = 2701 / 37 = 73
}

// Structs (구조체)
// A struct is a collection of fields.

type Vertex_1 struct {
	X int
	Y int
	// X, Y int라고 해도 동일함
}
func Practice3_2() {
	fmt.Println(Vertex_1{1, 2})
}

// Struct fields are accessed using a dot.

func Practice3_3() {
	v := Vertex_1{1, 2}
	v.X = 4
	fmt.Println(v)
}

// Pointers to structs
// Struct fields can be accessed through a struct pointer.
// 어떤 구조체에 pointer `p`가 있을 때, field인 `X`에 접근하려면 `(*p).X`라고 할 수 있다.
// 하지만 이렇게 하면 번거롭기 때문에 explicit dereference 없이 그냥 `p.X`라고 하는 것이 허용된다.

func Practice3_4() {
	v := Vertex_1{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

// Struct Literals
// A struct literal denotes a newly allocated struct value by listing the values of its fields
// You can list just a subset of fields by using the `Name:` syntax.
// And the order of named fields is irrelevant.
// The special prefix & returs a pointer to the struct value.

func Practice3_5() {
	var (
		v1 = Vertex_1{1, 2}   // has type Vertex_1
		v2 = Vertex_1{X: 1}   // Y: 0 is implicit
		v3 = Vertex_1{}       // X: 0 and Y: 0
		p  = &Vertex_1{10, 20}  // has type *Vertex
	)
	fmt.Println(v1, p, v2, v3)
	fmt.Println(p.X) // 10
}

// Arrays
// The type [n]T: `n`개의 type `T` 값을 가진 배열
func Practice3_6() {
	// Array의 length는 type에 포함되어 있음.
	var a [2]string
	fmt.Println(a)
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 아래는 Go의 Array literal syntax
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// Slices
func Practice3_7() {
	
}


func Practice3_8() {
	
}


func Practice3_9() {
	
}


func Practice3_10() {
	
}

func Practice3_11() {
	
}

func Practice3_12() {
	
}

func Practice3_13() {
	
}

func Practice3_14() {
	
}

func Practice3_15() {
	
}

func Practice3_16() {
	
}

func Practice3_17() {
	
}

func Practice3_18() {
	
}

func Practice3_19() {
	
}

func Practice3_20() {
	
}
