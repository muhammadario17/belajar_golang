package main

import "fmt"

func main() {
	//Variabel Biasa
	nama := "Ario"

	//Deklarasi Pointer
	var ptrNama *string = &nama

	fmt.Println("Alamat memori: ", ptrNama)
	fmt.Println("Isi data asli: ", *ptrNama)

	//Mengubah data asli lewat pointer
	*ptrNama = "Muhammad Ario"

	fmt.Println("Nama setelah diubah: ", nama)
}
