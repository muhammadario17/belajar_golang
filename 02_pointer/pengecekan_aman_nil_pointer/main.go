package main

import "fmt"

func cetakDiskon(persenDiskon *int) {
	//pengecekan keamanan
	if persenDiskon == nil {
		fmt.Println("Diskon: Tidak ada diskon (0%)")
		return
	}

	//Aman untuk di-dereference (*) karena sudah dicek bukan nil
	fmt.Printf("Diskon: Diskon sebesar %d%%\n", *persenDiskon)
}

func main() {
	var promo *int = nil //Belum ada diskon
	cetakDiskon(promo)   //Output: Diskon: Tidak ada diskon (0)

	nominalDiskon := 20
	promo = &nominalDiskon
	cetakDiskon(promo) //Output: Diskon: Diskon sebesar 20%
}
