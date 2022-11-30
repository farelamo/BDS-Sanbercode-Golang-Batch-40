package main

import (
	"fmt"
	"math"
	"strings"
)

/****************** Penerapan Interface ******************/
type hitung interface {
	luas() float64
	keliling() float64
}

// lingkaran
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// persegi
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

/****************** Embedded Interface ******************/
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitungEmbbed interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k kubus) keliling() float64 {
	return k.sisi * 12
}


func main(){

	/****************** Penerapan Interface ******************/
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("======== Persegi ========")
	fmt.Println("Luas", bangunDatar.luas())
	fmt.Println("Keliling", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("======= Lingkaran =======")
	fmt.Println("Luas", bangunDatar.luas())
	fmt.Println("Keliling", bangunDatar.keliling())
	fmt.Println("Jari-Jari", bangunDatar.(lingkaran).jariJari())

	/****************** Embedded Interface ******************/
	var bangunRuang hitungEmbbed
	
	bangunRuang = kubus{2}
	fmt.Println("======= Kubus =======")
	fmt.Println("Luas", bangunRuang.luas())
	fmt.Println("Keliling", bangunRuang.keliling())
	fmt.Println("Volume", bangunRuang.volume())

	fmt.Println()

	/****************** Interface Kosong ******************/
	// Type data yang sangat special, karna bisa menampung segala jenis tipe data, termasuk array, pointer, dll
	// Type data dengan konsep ini disebut dengan dynamic typing

	var secret interface{}

	secret = "inipassword456"
	fmt.Println(secret)

	secret = []string{"apel", "mangga", "jeruk"}
	fmt.Println(secret)

	secret = 9.11
	fmt.Println(secret)

	fmt.Println()
	fmt.Println("======= Casting Interface Kosong =======")
	fmt.Println()
	// Casting pada interface kosong
	// Hasil yang ditampilkan ke output oleh interface kosong adalah string (bukan nilai asli)
	// Untuk mengoutput nilai asli dalam interface kosong bisa seperti berikut ini

	var angka interface{}

	angka = 2
	var number = angka.(int) * 10
	fmt.Println(angka, "multipled by 10 is :", number)

	angka 		= []string{"Hani", "Tompel", "Koyek"}
	var result 	= strings.Join(angka.([]string), ", ") // var angka perlu dicasting dahulu ke []string
	fmt.Println(result, "is my friend call")

	// Teknik casting ini disebut Type Assertions

	fmt.Println()
	
	/****************** Casting Interface Kosong Ke Object Pointer ******************/
	type person struct {
		name string
		age int
	}

	var people interface{} = &person{name: "Helena", age: 20}
	var name = people.(*person).name
	fmt.Println(name)
}