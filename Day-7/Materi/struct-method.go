package main

import (
	"fmt"
)

func main(){

							/************** STRUCT **************/
	// 1. Deklarasi Struct 
	type student struct {
		name string
		grade int
	}

	// 2. Penerapan Struct
	var john student
	john.name  = "john doe"
	john.grade = 80

	fmt.Println("name", john.name)
	fmt.Println("grade", john.grade)

	// 3. Struct Literals
	// Beberapa cara membuat object dari struct

	// A. Cara pertama
	var wick   = student{}
	wick.name  = "wick"
	wick.grade = 90

	// B. Cara kedua
	var doe = student{"doe", 100}

	// C. Cara ketiga
	var jack = student{name: "jack", grade: 95}

	fmt.Println("Student 1 :", wick.name)
	fmt.Println("Student 2 :", doe.name)
	fmt.Println("Student 3 :", jack.name)

	// 4. Embedded Struct
	type person struct {
		name string
		age int
	}

	type teacher struct {
		special string
		person //struct person
	}

	// A. Cara Pertama
	var upin  	 	= teacher{}
	upin.name 	 	= "Upin Saifudin"
	upin.age  	 	= 30
	upin.special 	= "mathematics"
	upin.person.age = 28

	// B. Cara Kedua
	var ipinData = person{name: "Ipin Maemunah", age: 25}
	var ipin     = teacher{person: ipinData, special: "chemistry"}

	fmt.Println("name  :", ipin.name)
	fmt.Println("age   :", ipin.age)
	fmt.Println("grade :", ipin.special)
	
	// 5. Anonymous Struct
	// Tdk dideklarasikan di awal seperti nomor 1-4 sebagai tipe data baru, 
	// Bagus dipakai ketika pembuatan object struct yang hanya dipakai sekali saja

	type country struct {
		name string
	}

	var car = struct {
		merk string
		cc   int
		country
	}{}

	car.merk 	= "BMW"
	car.cc   	= 200
	car.country = country{name: "Indonesia"}

	fmt.Println("Car Name :", car.merk)
	fmt.Println("Car CC :", car.cc)
	fmt.Println("Used by :", car.country)
	

	// Aturan anonymous struct adalah penulisan object untuk inisialisasi setelah deklarasi
	// walau tanpa property / kosong, tetap harus ditulis

	// tanpa property
	var anjay = struct {
		sentence string
		limit int
	}{}

	anjay.limit = 5

	// dengan property
	var hehe = struct {
		sentence string
		min int
	}{
		sentence: "slebeww",
		min: 20,
	}

	fmt.Println(hehe)

	// 6. Nested Struct
	type bike struct {
		owner struct {
			name string
			age int
		}
		merk string
	}
}

					/************** METHOD **************/
	
// Perbedaan method dengan function biasa 
// 1. method punya akses ke property data struct hingga level private
// 2. Proses bisa dienkapsulasi dengan baik
// 3. Penulisan deklarasinya di luar func main

type planet struct {
	name string
	distance int
}

func (s planet) sayHello() {
	fmt.Println("Hello", s.name)
}

// func syaHello yang dideklarasi akan menjadi milik struct planet
// jadi bisa diakses lewat struct planet, planet.sayHello()
/*
	var sun = planet{"Sun", "1000000000"}
	sun.sayHello()
*/