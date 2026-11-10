package main

import "fmt"

type Mahasiswa struct {
	NIM     string
	Nama    string
	Jurusan string
	IPK     float64
}

func tampilkanSemua(daftar []Mahasiswa) {
	fmt.Printf("\n%-12s %-20s %-25s %s\n", "NIM", "NAMA", "JURUSAN", "IPK")
	fmt.Println("-------------------------------------------------------")
	for _, m := range daftar {
		fmt.Printf("%-12s %-20s %-25s %.2f\n", m.NIM, m.Nama, m.Jurusan, m.IPK)
	}
}

func cariMahasiswa(daftar []Mahasiswa, nimDicari string) *Mahasiswa {
	for i := range daftar {
		if daftar[i].NIM == nimDicari {
			return &daftar[i]
		}
	}
	return nil
}

func main() {
	fmt.Println("=====SISTEM DATABASE MAHASISWA (SLICE OF STRUCT)=====")

	daftarMhs := []Mahasiswa{
		{NIM: "17181920", Nama: "Amba", Jurusan: "Teknik Pengeboran", IPK: 3.80},
		{NIM: "17181921", Nama: "Imut", Jurusan: "Teknik Penghitaman", IPK: 3.81},
		{NIM: "17181922", Nama: "Rusdi", Jurusan: "Teknik Pertambangan", IPK: 3.85},
	}

	tampilkanSemua(daftarMhs)

	mhsBaru := Mahasiswa{NIM: "17181923", Nama: "Gatot", Jurusan: "Teknik Penghisapan", IPK: 3.50}
	daftarMhs = append(daftarMhs, mhsBaru)
	fmt.Println("\n [INFO] 1 Mahasiswa baru berhasil ditambahkan!")

	tampilkanSemua(daftarMhs)

	fmt.Println("\n---Fitur Pencarian---")
	nimCari := "17181923"
	hasil := cariMahasiswa(daftarMhs, nimCari)

	if hasil != nil {
		fmt.Printf("Mahasiswa ditemukan: %s | Jurusan: %s | IPK: %.2f\n",
			hasil.Nama, hasil.Jurusan, hasil.IPK)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan. \n", nimCari)
	}
}
