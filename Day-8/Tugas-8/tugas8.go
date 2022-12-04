package main

import (
	"fmt"
	"math"
	"strconv"
	// "strings"
)

/**************** Soal 1 ****************/
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas + s.tinggi
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * (float64(b.panjang)*float64(b.lebar) +
		float64(b.panjang)*float64(b.tinggi) +
		float64(b.lebar)*float64(b.tinggi))
}

/**************** Soal 2 ****************/
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type cetakData interface {
	print() string
}

func (p phone) print() string {
	var colors string
	for _, color := range p.colors {
		colors += color + ", "
	}
	return "name : " + p.name + "\nbrand : " + p.brand + "\nyear : " + strconv.Itoa(p.year) + "\ncolors : " + colors
}

func main() {
	// Soal 1
	var bangunDatar hitungBangunDatar

	bangunDatar = segitigaSamaSisi{2, 3}
	fmt.Println("==== Segitiga Sama Sisi ====")
	fmt.Println("Keliling :", bangunDatar.keliling())
	fmt.Println("Luas :", bangunDatar.luas())

	bangunDatar = persegiPanjang{5, 2}
	fmt.Println("==== Persegi Panjang ====")
	fmt.Println("Keliling :", bangunDatar.keliling())
	fmt.Println("Luas :", bangunDatar.luas())

	var bangunRuang hitungBangunRuang

	bangunRuang = tabung{7.0, 5.0}
	fmt.Println("==== Tabung ====")
	fmt.Println("Volume :", bangunRuang.volume())
	fmt.Println("Luas Permukaan :", bangunRuang.luasPermukaan())

	bangunRuang = balok{10, 5, 5}
	fmt.Println("==== Balok ====")
	fmt.Println("Volume :", bangunRuang.volume())
	fmt.Println("Luas Permukaan :", bangunRuang.luasPermukaan())

	fmt.Println() // untuk jarak soal

	// Soal 2
	var output cetakData

	output = phone{name: "Samsung Galaxy Note 20", brand: "Samsung", year: 2020, colors: []string{"merah", "biru"}}
	fmt.Println(output.print())

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

	fmt.Println()

	// Soal 4

	var prefix interface{} = "hasil penjumlahan dari"
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	pertama := kumpulanAngkaPertama.([]int)[0:1:1][0]
	kedua 	:= kumpulanAngkaPertama.([]int)[1:2:2][0]
	ketiga 	:= kumpulanAngkaKedua.([]int)[0:1:1][0]
	keempat := kumpulanAngkaKedua.([]int)[1:2:2][0]

	fmt.Println(prefix, pertama, "+", kedua, "+", ketiga, "+", keempat, "=", (pertama + kedua + ketiga + keempat))
	// bingung casting bang, dari yang interface slice int ke slice string
}
