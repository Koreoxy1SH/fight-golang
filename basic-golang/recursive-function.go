package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	//Recursive function adalah function yang memanggil function dirinya sendiri
	//kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
	//contoh kasus yang lebih mudah diselesaikan menggunkana recursive function adalah Factorial
	fmt.Println(factorialLoop(5))

	fmt.Println(factorialRecursive(10))
	//Recursive function = function yang memanggil dirinya sendiri
}
