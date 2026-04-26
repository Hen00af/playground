	// interfave

package main

import (
	"fmt"
)

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
	dist string
}

func (g Greeter) Talk() {
	fmt.Printf("my name is %s\n", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{dist: "Sendai",name: "hen00af"}
	talker.Talk()
}
