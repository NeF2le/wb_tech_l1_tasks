package main

import (
	"fmt"
	"reflect"
)

func defineType(x interface{}) {
	t := reflect.TypeOf(x)
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("chan")
	default:
		panic("unhandled default case")
	}
}

func main() {
	a := 3
	b := "hello world"
	c := false
	d := make(chan int)

	defineType(a)
	defineType(b)
	defineType(c)
	defineType(d)
}
