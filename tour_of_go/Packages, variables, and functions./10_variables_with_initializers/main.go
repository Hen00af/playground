package main

import "fmt"

var i, j int = 1, 2

func main() {
	// 変数は初期化される際の値によって型が自動的に割り振られる
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)	
}