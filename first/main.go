package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) PrintName() {
	fmt.Println(h.name)

}

func (h Human) PrintAge() {
	fmt.Println(h.age)
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
	action.PrintAge()
}
