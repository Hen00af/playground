package main

import "fmt"
// 返り値が同じだとひとまとめにできる
// (x int, y int) -> (x, y int)
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
