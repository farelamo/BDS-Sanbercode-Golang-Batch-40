package main

import (
	"fmt"
	"strconv"
	"Tugas-9/library"
)

func main() {
	// Soal 1
	var bangunDatar library.HitungBangunDatar

	bangunDatar = library.SegitigaSamaSisi{2, 3}
	fmt.Println("==== Segitiga Sama Sisi ====")
	fmt.Println("Keliling :", bangunDatar.Keliling())
	fmt.Println("Luas :", bangunDatar.Luas())

	bangunDatar = library.PersegiPanjang{5, 2}
	fmt.Println("==== Persegi Panjang ====")
	fmt.Println("Keliling :", bangunDatar.Keliling())
	fmt.Println("Luas :", bangunDatar.Luas())

	var bangunRuang library.HitungBangunRuang

	bangunRuang = library.Tabung{7.0, 5.0}
	fmt.Println("==== Tabung ====")
	fmt.Println("Volume :", bangunRuang.Volume())
	fmt.Println("Luas Permukaan :", bangunRuang.LuasPermukaan())

	bangunRuang = library.Balok{10, 5, 5}
	fmt.Println("==== Balok ====")
	fmt.Println("Volume :", bangunRuang.Volume())
	fmt.Println("Luas Permukaan :", bangunRuang.LuasPermukaan())

	fmt.Println() // untuk jarak soal

	// Soal 2
	var output library.CetakData

	output = library.Phone{Name: "Samsung Galaxy Note 20", Brand: "Samsung", Year: 2020, Colors: []string{"merah", "biru"}}
	fmt.Println(output.Print())

	fmt.Println() // untuk jarak soal

	// Soal 3
	var luasPersegi = func(sisi int, kondisi bool) interface{} {
		if kondisi {
			if sisi > 0 {
				sisi := sisi * sisi
				return "Luas persegi dengan sisi 2 cm adalah " + strconv.Itoa(sisi) + " cm"
			} else {
				return "Maaf anda belum menginput sisi dari persegi"
			}
		} else {
			if sisi > 0 {
				return strconv.Itoa(sisi)
			} else {
				return "nil"
			}
		}
	}

	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// Soal 4

	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	fmt.Println(prefix, kumpulanAngkaPertama, kumpulanAngkaKedua)
	// bingung casting bang, dari yang interface slice int ke slice string
}
