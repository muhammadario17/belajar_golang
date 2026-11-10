package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mahasiswa struct {
	NIM     string
	Nama    string
	Jurusan string
	IPK     float64
}

var scanner = bufio.NewScanner(os.Stdin)

func inputTeks(label string) string {
	fmt.Print(label)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func tampilanDaftar(daftar []Mahasiswa) {
	if len(daftar) == 0 {
		fmt.Println("\n Belum ada data mahasiswa!")
		return
	}
	fmt.Printf("\n%-12s %-20s %25s %s\n", "NIM", "NAMA", "JURUSAN", "IPK")
	fmt.Println("----------------------------------------------------")
	for _, m := range daftar {
		fmt.Printf("%-12s %-20s %-25s %.2f\n", m.NIM, m.Nama, m.Jurusan, m.IPK)
	}
}

func tambahMahasiswa(daftar *[]Mahasiswa) {
	fmt.Println("\n---Form Tambah Mahasiswa---")
	nim := inputTeks("Masukan NIM	: ")
	nama := inputTeks("Masukan Nama	: ")
	jurusan := inputTeks("Masukan Jurusan	: ")
	ipkStr := inputTeks("Masukan IPK	: ")

	ipk, err := strconv.ParseFloat(ipkStr, 64)
	if err != nil {
		fmt.Println("Gagal: format IPK harus berupa angka desimal (Contoh: 3.00)")
		return
	}

	mhsBaru := Mahasiswa{
		NIM:     nim,
		Nama:    nama,
		Jurusan: jurusan,
		IPK:     ipk,
	}
	*daftar = append(*daftar, mhsBaru)
	fmt.Println("Mahasiswa berhasil disimpan")
}

func cariMahasiswa(daftar []Mahasiswa) {
	fmt.Println("\n---Fitur Pencarian---")
	nimCari := inputTeks("Masukan NIM yang dicari:	")

	for _, m := range daftar {
		if m.NIM == nimCari {
			fmt.Printf("Ditemukan: %s | Jurusan: %s | IPK: %.2f\n",
				m.Nama, m.Jurusan, m.IPK)
			return
		}
	}
	fmt.Printf("Mahasiswa dengan NIM '%s' tidak ditemukan.\n", nimCari)
}

func main() {
	databaseMhs := []Mahasiswa{
		{NIM: "17181920", Nama: "Amba", Jurusan: "Teknik Pengeboran", IPK: 3.80},
		{NIM: "17181921", Nama: "Imut", Jurusan: "Teknik Penghitaman", IPK: 3.81},
	}

	for {
		fmt.Println("\n========================================")
		fmt.Println("🎓 SISTEM MANAJEMEN MAHASISWA (CLI)")
		fmt.Println("========================================")
		fmt.Println("1. Lihat Semua Mahasiswa")
		fmt.Println("2. Tambah Mahasiswa Baru")
		fmt.Println("3. Cari Mahasiswa via NIM")
		fmt.Println("4. Keluar")
		pilihan := inputTeks("Pilih menu [1-4]: ")

		switch pilihan {
		case "1":
			tampilanDaftar(databaseMhs)
		case "2":
			tambahMahasiswa(&databaseMhs)
		case "3":
			cariMahasiswa(databaseMhs)
		case "4":
			fmt.Println("\nTerima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silahkan pihih 1-4")
		}
	}
}
