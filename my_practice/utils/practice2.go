// Flow control statements: for, if, else, switch, and defer

package utils

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// Note: Unlike other languages like C, Java, or JavaScrip
// there are no parentheses surrounding the three components of the for statement,
// and the braces { } are always required.

func Practice2_1() {
// the init statement: executed before the first iteration
// the condition expression: evaluated before every iteration
// the post statement: executed at the end of every iteration

// The init statement will often be a short variable declaration,
// and the variables declared there are visible only in the scope of the for statement.

	sum := 0
	for i:=0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func Practice2_2() {
	// The init and post statements are optional.
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func Practice2_3() {
	// for is Go's "while"
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func Practice2_4() {
	// Omitting the loop condition makes an infinite loop
	// for {
	// }
}

func sqrt2_5(x float64) string {
	// Similar to for loops, in if statements, the expression need not be
	// surrounded by parentheses () but the braces {} are required
	if x < 0 {
		return sqrt2_5(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Practice2_5() {
	fmt.Println(sqrt2_5(2), sqrt2_5(-4))	
}

func pow2_6(x, n, lim float64) float64 {
	// Like for, the if statement can start with a short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim // if a use v here, it causes an error
}

func Practice2_6() {
	fmt.Println(
		pow2_6(3, 2, 10),
		pow2_6(3, 3, 20),
	)
}

func pow2_7(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
		// can't use v here, though
	}
	return lim
}

func Practice2_7() {
	fmt.Println(
		pow2_7(3, 2, 10),
		pow2_7(3, 3, 20),
	)
}

func sqrt2_8(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}


func ImprovedSqrt(x float64) float64 {
	z := 1.0
	epsilon := 1e-10 // very small value to check for changes
	for {
		nextZ := z - (z*z-x)/(2*z)
		if math.Abs(nextZ-z) < epsilon {
			break
		}
		z = nextZ
	}
	return z
}

func Practice2_8() {
	fmt.Println(sqrt2_8(2))
	fmt.Println(ImprovedSqrt(2))
}

// Go only runs the selected case, not all the cases that follow.
// Thus, the `break` statement that is needed at the end of each case is provided automatically in Go.
// Another important difference is that Go's switch cases need not be constants,
// and the values involved need not be integers.

func Practice2_9() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux" :
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s. \n", os)
	}
}

func Practice2_10() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0: // today + 0 == time.Saturday
		fmt.Println("Today!")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// Switch without a condition is the same as switch true.
func Practice2_11() {
	t:=time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning.")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// A defer statement defers the execution of a function
// until the surrounding function returns.
// The deferred call's arguments are evaluated immediately,
// but the function call is not executed until the surrounding function returns.

func Practice2_12() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// Deferred function calls are pushed onto a stack. When a function returns,
// its deferred calls are executed in last-in-first-out order.
func Practice2_13() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
