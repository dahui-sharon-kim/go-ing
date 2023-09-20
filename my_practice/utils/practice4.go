package utils

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
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

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Practice4_3() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
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
// Functions that take a value argument must take a value of that specific type.
// Methods with value receivers take either a value or a pointer as the receiver when they are called.

func Practice4_7() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4,3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

// Choosing a value or pointer receiver
// There are two reasons to use a pointer receiver.
// The first is so that the method can modify the value that its receiver points to.
// The second is to avoid copying the value on each method call.
// This can be more efficient if the receiver is a large struct, for example.

// In this example, both Scale and Abs are methods with receiver type *Vertex,
// even though the Abs method needn't modify its receiver.
// In general, all methods on a given type should have either value or pointer receivers,
// but not a mixture of both. (We'll see why over the next few pages.)

func Practice4_8() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

// Interfaces

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
// 인터페이스 타입을 가지는 어떠한 값은 인터페이스 내의 메소드를 구현하는 어떠한 값이든 가질 수 있다.

type Abser interface {
	Abs() float64
}

func Practice4_9() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f   // a MyFloat implements Abser
	fmt.Println(a.Abs()) // 1.4142135623730951
	a = &v  // a *Vertex implements Abser
	fmt.Println(a.Abs()) // 5
}

// Interfaces are implemented implicitly

// 타입은 메소드를 구현함으로써 인터페이스를 구현한다. A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword.
// Implicit interfaces decouple the definition of an interface from its implementation,
// which could then appear in any package without prearrangement.

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.

func (t T) M() {
	fmt.Println(t.S)
}

func Practice4_10() {
	var i I = T{"hello"}
	i.M()
}

// Interface values (== instance)

// Under the hood, interface values can be thought of a s a tuple of a value and a concrete type: (value, type)
// An interface value holds a value of a specific underlying concrete type.
// Calling a method on an interface value executes the method of the same name on its underlying type.

type I2 interface {
	M2()
}

func (t *T) M2() {
	fmt.Println(t.S)
}

type F float64

func (f F) M2() {
	fmt.Println(f)
}

func Practice4_11() {
	var i I2
	i = &T{"hello"} // interface value == instance
	describe(i)     // (&{hello}, *utils.T)
	i.M2()          // hello

	i = F(math.Pi)  // interface value == instance  
	describe(i)     // (3.141592653589793, utils.F)
	i.M2()          // 3.141592653589793
}

func describe(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}


// Interface values with nil underlying values

// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
// In some languages this would trigger a null pointer exception, but in Go it is common to write methods
// that gracefully handle being called with a nil receiver
// (as with the method `M` in this example)

// Note that an interface value that holds a nil concrete value is itself non-nil.

type I3 interface {
	M3()
}

func (t *T) M3() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S) // t
}

func Practice4_12() {
	var i I3

	var t *T
	i = t
	describe3(i)     // (<nil>, *utils.T)
	i.M3()           // <nil>

	i = &T{"hello"} 
	describe3(i)     // (&{hello}, *utils.T)
	i.M3()           // hello
}

// Nil interface values

// A nil interface value holds neither value nor concrete type.
// Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple
// to indicate which concrete method to call.

func describe3(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Nil interface values

// A nil interface value holds neithe value nor concrete type.
// Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to
// indicate which concrete method to call.

func Practice4_13() {
	var i I3
	describe3(i)
	i.M3() // run-time error
}

// The empty interface

// The interface type that specifies zero methods is known as the empty interface: `interface {}`
// An empty interface may hold values of any type. (Every type implements at least zero methods.)
// Empty interfaces are used by code that handles values of unknown type.
// For example, `fmt.Print` takes any number of arguments of type `interface{}`


func describe4(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func Practice4_14() {
	var i interface{}
	describe4(i)

	i = 42
	describe4(i)

	i = "hello"
	describe4(i)
}


// Type assertions (타입 단언)

// A type assertion provides access to an interface value's underlying concrete value.
// `t := i.(T)`
// This statement asserts the interface value i holds the concrete type T and
// assigns the underlying T value to the variable t.

// If `i` does not hold a `T`, the statement will trigger a panic.
// To test whether an interface value holds a psecific type, a type assertion can return two values:
// the underlying valud and a boolean value that reports whethr the assertion succeeded.

// `t, ok := i.(T)`
// If `i` holds a `T`, then `t` will be the underlying value and `ok` will be true.
// If not, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.
// Note the similarity btw this syntax and that of reading from a map.

func Practice4_15() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true ("hello" is the underlying value)

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false (float64 없기 때문에 float64의 zero value를 print)

	f = i.(float64) // panic
	fmt.Println(f)
}


// Type switches

// A type switch is a construct that permits several type assertions in series.
// A type switch is like a regular switch statement, but the cases in a type switch specify types (not values),
// and those values are compared against the type of the value held by the given interface value.

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// The declaration in a type switch has the same syntax as a type assertion `i.(T)`, but the specific type `T` is replaced with the keyword `type`.
// This switch statement tests whether the interface value `i` holds a value of type of `int` or `string` (in the example)

func Practice4_16() {
	do(21)
	do("hello")
	do(true)
}

// Stringers

// One of the most ubiquitos interfaces is Stringer defined by the fmt package. A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

type Stringer interface {
	String() string
} 

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func Practice4_17() {
	a := Person{"Arthur Dent", 42}	
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func Practice4_18() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// Errors

// Go programs express error state with `error` values.
// The `error` type is a built-in interface similar to `fmt.Stringer`:
// type error interface {
// 		Error() string
//}
// (As with `fmt.Stringer`, the `fmt` package looks for the `error` interface when printing values.)
// Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.
// A nil `error` denotes success; a non-nil `error` denotes failure

// Defined a custom error type called MyError
type MyError struct {
	When time.Time
	What string
}

// Implemented `Error` method for `MyError` type
// In Go, to make a type an error, you need to implement the `error` interface, which consists of a single method called `Error() string`. This method returns a string representation of the error.

func (e *MyError) Error() string {
	return fmt.Sprintf("%s (at %v)", e.What, e.When)
}

// Define `run` function, which returns an error.
func run() error {
	// Create an instance of `MyError`
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// In the context of error handling:
// When a function returns `nil` as an error, it typically means that no error occurred, and the operation was successful
// When a function returns non-`nil` error, it means an error did occur, and the error object contains information about what went wrong.
func Practice4_19() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// Exercise: Errors

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	}
	return ""
}

func Sqrt4_20(x float64) (float64, error) {
	if x < 0 {
		// ErrNegativeSqrt(x) is not calling a function.
		// Instead, it's creating a new value of type ErrNegativeSqrt with the value x
		// This works because ErrNegativeSqrt is an alias for float64
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func Practice4_20() {
	fmt.Println(Sqrt4_20(2))
	fmt.Println(Sqrt4_20(-2))
}

// Readers

// The io package specifies the io.Reader interface, which represents the read end of a stream of data. The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.

// The io.Reader interface has a Read method:
// `func (T) Read(b []byte) (n int, err error)`
// Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.

// The example code creates a strings.Reader and consumes its output 8 bytes at a time.

func Practice4_21() {
	// strings.NewReader() 함수는 `*strings.Reader` 타입의 값을 만들고 반환
	// 이 타입은 `io.Reader` 인터페이스를 구현
	r := strings.NewReader("Hello, Reader!") // 변수 r은 *strings.Reader 타입의 값.
	// strings.NewReader("Hello, Reader") creates a Reader that reads from the provided string.
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// Exercise: Readers

type MyReader4_22 struct {}

func (r MyReader4_22) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func Practice4_22() {
	// reader.Validate(MyReader4_22{})
}

// Exercise: rot13Reader

// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

type rot13Reader struct {
	r io.Reader
}

func Practice4_23() {
	// s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// r := rot13Reader{s}
	// io.Copy(os.Stdout, &r)
}

func Practice4_24() {
	
}

func Practice4_25() {
	
}

func Practice4_26() {
	
}