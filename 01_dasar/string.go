package main

import "fmt"

func main() {
	//Contoh tipe data string
	fmt.Println("Muhammad")
	fmt.Println("Muhammad Ario")
	fmt.Println("Muhammad Ario Luthfianto")

	//Penggunaan function len untuk menghitung jumlah karakter
	fmt.Println(len("Muhammad"))
	fmt.Println(len("Muhammad Ario"))
	fmt.Println(len("Muhammad Ario Luthfianto"))

	//Mengambil karakter pada posisi yang ditentukan oleh []
	fmt.Println("Muhammad"[0])
	fmt.Println("Muhammad Ario"[1])
	fmt.Println("Muhammad Ario Luthfianto"[2])
}
