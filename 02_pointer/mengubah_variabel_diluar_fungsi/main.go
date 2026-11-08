package main

import "fmt"

//Parameter menerima *int (alamat memori angka)
func naikLevel(levelSekarang *int) {
	*levelSekarang = *levelSekarang + 1 //Menambah nilai di alamat aslinya
}

func main() {
	levelAkun := 5

	//kirim alamat memorinya menggunakan &
	naikLevel(&levelAkun)

	fmt.Println("Level akun sekarang: ", levelAkun)
}
