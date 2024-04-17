//go:build ignore

package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}
func (t *T) N() {
	fmt.Println("happy")
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	t := T{"hello"}

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	t.M()
	t.N()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
