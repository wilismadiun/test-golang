package main

import "fmt"

// Tambah menjumlahkan dua bilangan integer
func Tambah(a, b int) int {
	return a + b
}

func main() {
	hasil := Tambah(10, 20)
	fmt.Println("Hasil:", hasil)
}
