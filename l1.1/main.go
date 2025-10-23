package wbl1

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) PrintName() {
	fmt.Println(h.name)

}

type Action struct {
	Human
}
