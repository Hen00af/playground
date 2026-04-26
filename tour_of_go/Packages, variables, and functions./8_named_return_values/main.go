package main

import "fmt"
// goでの戻り値となる変数に名前をつけることができます。
// 戻り値に名前をつけると、関数の最初で定義した変数名として扱われます。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
