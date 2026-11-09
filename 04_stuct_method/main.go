package main

import "fmt"

//1. Mendefinisikan STRUCT
type mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
	IPK     float64
	IsLulus bool
}

// 2. Menempelkan METHOD (Perilaku): Cetak Profil Mahasiswa
// (m Mahasiswa) disebut Value Receiver karena hanya membaca data saja
func (m mahasiswa) KenalkanDiri() {
	fmt.Printf("Hallo! saya %s (NIM: %s), jurusan %s. IPK saat ini: %.2f \n",
		m.Nama, m.NIM, m.Jurusan, m.IPK)
}

// 3. Menempelkan METHOD dengan POINTER: Update Nilai IPK
// Wajib pakai (m *Mahasiswa) agar nilai IPK fisik di memori benar-benar ter-update!
func (m *mahasiswa) UpdateIPK(ipkBaru float64) {
	m.IPK = ipkBaru
	if m.IPK >= 3.00 {
		m.IsLulus = true
	}
	fmt.Printf("--> IPK %s berhasil diperbarui menjadi: %.2f\n", m.Nama, m.IPK)
}

func main() {
	fmt.Println("===SISTEM AKADEMIK KAMPUS===")

	mhs1 := mahasiswa{
		Nama:    "M. Ario",
		NIM:     "15161780",
		Jurusan: "Teknik Penghitaman",
		IPK:     3.45,
		IsLulus: false,
	}

	mhs1.KenalkanDiri()

	fmt.Println("\nSetelah Ujian Semester Berakhir...")
	mhs1.UpdateIPK(3.83)

	fmt.Println("\nCek Status Kelulusan :")
	fmt.Printf("Nama %s | IPK Akhir %.2f | Status lulus: %t\n",
		mhs1.Nama, mhs1.IPK, mhs1.IsLulus)
}
