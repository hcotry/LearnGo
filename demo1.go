package main

import (
	"fmt"
)

func main() {
	var a string = "hello world!"
	fmt.Println(a)

	var b, c int = 1, 2

	fmt.Println(b, c)

	/*
			intVal := 1 

		var intVal int
		intVal =1
	*/

	intVal := 1
	fmt.Println(intVal)

	g, h := 123, "hello"

	fmt.Println(g, h)

	var age = 41
	age = age + 1
	age += 1
	age++

}
