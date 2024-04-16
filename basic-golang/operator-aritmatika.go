package main

import "fmt"

func main() {

	var a = 20
	var b = 30
	var d = 2
	var e = 2
	var c = a + b - d*e

	fmt.Println(c)

	//augmentend assingments
	var x = 10

	x += 10
	x -= 10
	x *= 20

	fmt.Println(x)

	//unary operator
	var i = 10
	i++ //i = 10 + 1
	i-- //j = 10 - 1

	fmt.Println(i)

}
