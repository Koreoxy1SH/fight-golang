package helper

var version = "1.0.0"
var App = "golang"

// FUNCTION INI TIDAK BISA DIAKSES DENGAN PACKAGE YANG BERBDA HANYA BISA DENGAN PACKAGE YANG SAMA,
// KARENA MENGGUNAKAN NAMA HURUF KECIL DIAWAL PENAMAANNYA
func sayBye(name string) string {
	return "Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
