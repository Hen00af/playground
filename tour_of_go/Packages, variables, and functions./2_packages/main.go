// Goはパッケージで構成される
//　プログラムはmainパッケージで開始される。
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is ", rand.Intn(10))
}
