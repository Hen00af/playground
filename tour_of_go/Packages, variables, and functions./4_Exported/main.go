package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) // OK: Pi は exported
	// fmt.Println(math.pi) // NG: pi は unexported
}