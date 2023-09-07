package main

import (
	"go-study/my_practice/utils"
)

// func main() {
// 	fmt.Println("Hello, World!")
// }

// func main() {
// 	// format
// 	// 6자리 미만은 그대로 표시하고, 6자리가 넘어가면 지수로 표시
// 	fmt.Printf("Now you have %g problems.\n", math.Sqrt(1000000000000))
// }

// func main() {
// 	fmt.Println(math.Pi)
// }

// // type이 있어야 함
// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }

// // When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
// func add(x, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }

// // A function can return any number of results.
// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func main() {
// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)
// }

// named return values
// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

// func main() {
// 	fmt.Println(split(17))
// }

// // A var statement can be at package or function level. We see both in this example.
// var c, python, java bool

// func main() {
// 	var i int
// 	fmt.Println(i, c, python, java)
// }

func main() {
	utils.Practice1_1()
}
