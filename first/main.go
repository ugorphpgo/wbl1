package main

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

func main() {
	action := Action{
		Human{
			name: "cherepacha",
			age:  22,
		},
	}
	action.PrintName()
}
