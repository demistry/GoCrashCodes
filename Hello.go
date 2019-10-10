package main

import "fmt"

func main() {
	fmt.Printf("Hello world, welcome to my Go\n")
	//statically typed, you must use all variables
	// var age int = 40
	// var newFloat float32 = 1.64
	var newString = "This is my new string"
	var stringLength int = len(newString)
	const pi float32 = 3.142
	fmt.Println("String length is "+newString+" with length ", stringLength)

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
