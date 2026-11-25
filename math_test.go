package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambah(t *testing.T) {
	// Inisialisasi assert
	assert := assert.New(t)

	// Skenario pengujian
	hasil := Tambah(2, 3)
	harapan := 5

	// Menggunakan assert.Equal untuk membandingkan nilai
	assert.Equal(harapan, hasil, "Seharusnya hasil penjumlahannya adalah 5")

	// Contoh pengujian lain (misal: memastikan tidak bernilai nil/kosong, dll)
	assert.NotNil(hasil)
}
