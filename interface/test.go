//go:build ignore

package main

import "fmt"

type People struct {
	age  int
	name string
}

func (p *People) addAge() {
	p.age++
	fmt.Println(p.age)
}

func main() {
	jeff := People{19, "jeff"}

	jeff.addAge()

	fmt.Println(jeff.age)
}
