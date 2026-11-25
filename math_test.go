package main

import "testing"

func TestTambah(T *testing.T) {
	// Menyiapkan skenario pengujian
	hasil := Tambah(2, 3)
	harapan := 5

	if hasil != harapan {
		T.Errorf("Tambah(2, 3) menghasilkan %d, diharapkan %d", hasil, harapan)
	}
}
