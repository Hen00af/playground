package main

import "fmt"

// var statementは変数を宣言します。　関数の引数リストと同様に、複数の変数の最後に型を描くことで変数のリストを宣言することができます。

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}