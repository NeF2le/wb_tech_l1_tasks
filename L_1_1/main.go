package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	Phone string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.Phone)
}

type Employee struct {
	Human
	Company   string
	WorkEmail string
}

func main() {
	e := Employee{
		Human: Human{
			Name:  "James",
			Age:   32,
			Phone: "+1 (22) 33-22",
		},
		Company:   "WB",
		WorkEmail: "wbtech@gmail.com",
	}
	e.SayHi()
}
