package main

import (
	"fmt"
)

type buah struct {
	nama, warna string
	adaBijinya 	bool
	harga 		int
}

// Soal 2
type segitiga struct {
	alas,tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (L segitiga) luasSegitiga(){
	fmt.Println("Luas Segitiga", L.alas * L.tinggi)
}

func (L persegi) luasPersegi(){
	fmt.Println("Luas Persegi", L.sisi * L.sisi)
}


func (L persegiPanjang) luasPersegiPanjang(){
	fmt.Println("Luas Persegi Panjang", L.panjang * L.lebar)
}


// Soal 3 
type phone struct {
	name, brand string
	year int
	colors []string
}

func (p *phone) addColor(color string){
	p.colors = append(p.colors, color)
}

// Soal 4
type movie struct {
	title, genre string
	duration, year int
}

func tambahDataFilm(name string, duration int, genre string, year int, dataFilm *[]movie) {
	var addMovie = movie{title: name, genre: genre, duration: duration, year: year}
	*dataFilm    = append(*dataFilm, addMovie)
}

func main(){

	// Soal 1
	var nanas 		= buah{"Nanas", "Kuning", false, 9000}
	var jeruk 		= buah{nama: "Jeruk", warna: "Oranye", adaBijinya: true, harga: 8000}
	var semangka	= buah{}
		semangka.nama		= "Semangka" 
		semangka.warna		= "Hijau & Merah"
		semangka.adaBijinya	= true
		semangka.harga  	= 10000
	var pisang 		= buah{"Pisang", "Kuning", false, 5000}

	fmt.Println(nanas)
	fmt.Println(semangka)
	fmt.Println(jeruk)
	fmt.Println(pisang)

	fmt.Println() // untuk space di ouput terminal

	// Soal 2
	var segitiga = segitiga{alas: 3, tinggi: 4}
	segitiga.luasSegitiga()

	var persegi  = persegi{sisi: 4}
	persegi.luasPersegi()

	var persegiPanjang = persegiPanjang{panjang: 6, lebar: 4}
	persegiPanjang.luasPersegiPanjang()

	fmt.Println()  // untuk space di ouput terminal

	// Soal 3
	var phone = phone{
		name : "Samsung Galaxy A03S", 
		brand: "Samsung",
		year : 2022,
	}

	phone.addColor("merah")
	fmt.Println(phone)

	fmt.Println()  // untuk space di ouput terminal

	// Soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	fmt.Println(dataFilm)
}