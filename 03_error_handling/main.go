package main

import (
	"errors"
	"fmt"
)

// Fungsi tarikTunai mengembalikan DUA nilai sekaligus: (int, error)
func tarikTunai(saldo int, nominal int) (int, error) {
	//Validasi nominal tidak boleh minus atau nol
	if nominal <= 0 {
		return saldo, errors.New("Nominal penarikan harus lebih dari 0")
	}

	//Validasi apakah saldo mencukupi
	if nominal > saldo {
		return saldo, errors.New("Saldo tidak mencukupi untuk penarikan ini")
	}

	//Jika semua aman: kurangi saldo, dan errors nya bernilai nil (tidak ada error)
	sisaSaldo := saldo - nominal
	return sisaSaldo, nil
}

func main() {
	fmt.Println("===Simulasi Mesin ATM===")
	saldoKu := 100000

	//Percobaan tarik tunai
	fmt.Println("\n1. Mencoba tarik Rp.40.000...")
	sisa, err := tarikTunai(saldoKu, 40000)

	//Pola standar Go: Cek apakah ada yang error
	if err != nil {
		fmt.Println("Gagal", err)
	} else {
		saldoKu = sisa
		fmt.Println("Berhasil! Sisa saldo sekarang: Rp.", saldoKu)
	}

	//Percobaan tarik tunai melebihi saldo
	fmt.Println("\n2. Mencoba tarik Rp.200.000")
	sisa, err = tarikTunai(saldoKu, 200000)

	if err != nil {
		fmt.Println("Gagal !!!", err)
	} else {
		saldoKu = sisa
		fmt.Println("Berhasil! sisa saldo sekarang: Rp.", saldoKu)
	}

	fmt.Println("\nProgram selesai berjalan dengan aman tanpa crash")
}
