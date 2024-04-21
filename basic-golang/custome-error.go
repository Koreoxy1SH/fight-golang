//Membuat custome error
//karena error adalah sebuah interface, jadi jika ingin membuat error sendiri, kita bisa membuat struct yang mengikuti kontrak dari interface error

package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}

	if id != "budi" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {

	err := SaveData("s", nil)
	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error:", validationErr.Message)
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error:", notFoundErr.Message)
		} else {
			fmt.Println("Unknown error:", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
