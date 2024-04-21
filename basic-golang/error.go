//ERROR INTERFACE
//Golang memiliki interface yg digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error

//Membuat Error
//untuk membuat error, kita tidak perlu manual.
//golang sudah menyediakan library untuk membuat helper secara mudah, yg terdapat di package errors

package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}

}

func main() {

	result, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil ", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}
