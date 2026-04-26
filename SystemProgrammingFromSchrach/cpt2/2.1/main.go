package main

import (
	"fmt"
	"os"
)
func main() {
	fmt.Fprintf(os.Stdout, "1 + 1 = %2f\n", 1.1+1)
}