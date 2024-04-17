package main

import "fmt"

func main() {
	// Tipe data slice adalah potongan dari data array
	// Slice mirip dengan array, yang membedakan adalah ukuran Slice bisa berubah
	// Slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di array

	// Detail tipe data Slice
	// =======================
	// Tipe data slice memiliki 3 data, yaitu pointer, length dan capacity
	// Pointer adalah penunjuk data pertama di array para slice
	// Length adalah panjang dari slice, dan
	// capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

	name := [...]string{
		"brian",
		"dudi",
		"kuya",
		"huhu",
		"yuyu",
		"edo",
	}

	slice := name[4:6]
	slice2 := name[:6]
	slice3 := name[:2]
	slice4 := name[3:]
	slice5 := name[:]
	var slice6 []string = name[:]

	fmt.Println(name)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(append(slice, "TEST"))

	fmt.Println("====================")
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)

	fmt.Println("=======================")

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	sliceDays := days[5:]

	sliceDays[0] = "Sabtu baru"
	sliceDays[1] = "Minggu baru"
	fmt.Println(sliceDays)
	fmt.Println(days)

	daySlice2 := append(sliceDays, "libur baru")
	fmt.Println(daySlice2)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "dudi"
	newSlice[1] = "juki"
	//newSlice[2] = "juki" tidak bisak menambahkan array seperti ini harus menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "kiki")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[1] = "Luli"
	fmt.Println(newSlice2)

	//COPY DATA Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//PERBEDAAN ARRAY DAN SLICE

	iniArray := [...]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
