package utils

import (
	"fmt"
	"strings"
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
// 타입 [n]T: `n`개의 type `T` 값을 가진 배열
func Practice3_6() {
	// Array의 length는 type에 포함되어 있음.
	var a [2]string
	var b [3]float64
	fmt.Println(a)
	fmt.Println(b)
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 아래는 Go의 Array literal syntax
	primes := [6]int{2, 3, 5, 7, 11, 13} // **Initialized with values**
	fmt.Println(primes) // [2 3 5 7 11 13]
	var y [5]int = [5]int{10, 20, 30} // Partial assignment
	fmt.Println(y) // [10 20 30 0 0]
}

// Slices
func Practice3_7() {
	primes := [6]int{2, 3, 5, 7, 11, 13} 

	//타입 []T: T타입의 값을 가진 슬라이스
	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]
}

// Slices are like references to arrays

// A slice does not store any data, it just describes a section of an underlying array.
// Changing the elements of a slice modifies the corresponding elements of its underlying array.
// Other slices that share the same underlying array will see those changes.
func Practice3_8() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // 슬라이스의 element를 바꾸면 원본 array의 해당 element가 바뀐다
	fmt.Println(a, b)
	fmt.Println(names) // [John XXX George Ringo]
}

// Slice literals
// A slice literal is an array literal without the length

// [3]bool{true, true, false} // array literal
// []bool{true, true, false} // slice literal 
// ** slice literal은 array literal로 만드는 것과 똑같은 array를 만든 후
// ** 그 array를 reference하는 슬라이스를 build
func Practice3_9() {
	q := []int{2, 3, 5, 7, 11, 13} // 이때 array의 length가 6으로 고정됨
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	
}


func Practice3_10() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[:]
	fmt.Println(s)
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
}

// Slice length and capacity

// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the **underlying array**,
// **counting from the first element in the slice.**

// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.


func Practice3_11() {
	// 여기서 s는 array가 아니고 underlying array를 가리키는 slice
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice3_11(s)

	s = s[:]
	printSlice3_11(s) // len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0] // Slice the slice to give it zero length.
	printSlice3_11(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice3_11(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice3_11(s) // len=2 cap=4 [5 7]
}

// fyi: Forward declaration
// (같은 패키지 내에서는 선언하기 전에 사용할 수 있다)
func printSlice3_11(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Nil slices
func Practice3_12() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // [] 0 0
	if s == nil {
		fmt.Println("nil!")
	}
	// s[0] = 3 // 이 코드를 실행하면 에러 발생
}

// Creating a slice with make

// Slices can be created with the build-in `make` function
// This is how dynamically-sized arrays are made

func Practice3_13() {
	// The `make` function allocates a zeroed ARRAY and returns a SLICE
	// that refers to that array
	a := make([]int, 5)
	printSlice3_13("a", a) // a len=5 cap=5 [0 0 0 0 0]

	// To specify a capacity, pass a third argument to `make``:
	b := make([]int, 0, 5) // [0 0 0 0 0]라는 array을 할당한 후 []라는 slice를 반환
	printSlice3_13("b", b) // b len=0 cap=5 []

	c := b[:2]             // [0 0 0 0 0]이라는 array에서 [:2]만큼 슬라이스
	printSlice3_13("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]            // [0 0 0 0 0]이라는 array에서 [2:5]만큼 슬라이스 
	printSlice3_13("d", d) // d len=3 cap=3 [0 0 0]
	
}

func printSlice3_13(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}


// Slice of Slices

// Slices can contain any type, including other slices.
func Practice3_14() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}	
}

// Appending to a slice
// It is common to append new elements to a slice.

func Practice3_15() {
	var s []int
	printSlice3_15(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice3_15(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice3_15(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice3_15(s)
	
}

func printSlice3_15(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
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
