package main

import "fmt"

func main() {
	fmt.Println("Halo Dunia IT! Saya resmi mulai belajar Go & Git")

	// Pembuktian konsep memori (Fotokopi vs Link Google Docs):
	// 1. Tipe Primitif (Pass by Value / Fotokopi)
	angka1 := 10
	angka2 := angka1
	angka2 = 25

	fmt.Printf("\n---Tipe Primitif---\n")
	fmt.Println("Angka 1 :", angka1) //Tetap 10
	fmt.Println("Angka 2 :", angka2) //Berubah Jadi 25

	// 2. Slice/List (Pass by Reference / Link Google Docs)
	daftarA := []int{1, 2, 3}
	daftarB := daftarA

	daftarB[0] = 999

	fmt.Printf("\n---Tipe Slice/list (Alamat memori)---\n")
	fmt.Println("Daftar A:", daftarA)
	fmt.Println("Daftar B", daftarB)
}
