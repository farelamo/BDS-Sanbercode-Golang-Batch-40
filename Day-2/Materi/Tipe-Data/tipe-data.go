package main

import (
	"fmt"
	"strings"
)

func main() {

	/* 1. Numerik non decimal */
	var positifNumber uint8 = 100
	var negatifNumber = -123

	fmt.Printf("Bilangan Positif: %d\n", positifNumber)
	fmt.Printf("Bilangan Negatif: %d\n", negatifNumber)

	/* 2. Numerik Decimal */
	var decimalNumber = 2.62

	fmt.Printf("Bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("Bilangan desimal: %.3f\n", decimalNumber)

	// template %f mengubah float jadi string
	// template %.3f mengatur jumah angka dibelakang koma

	/* 3. Boolean */
	var isActive bool = true
	fmt.Printf("is Active ? %t \n:", isActive)

	// template %t u/ memformat data bool menggunakan fungsi fmt.Printf()

	/* 4. String */
	var message = `
		Nama saya "John Wick". 
		Salam kenal. 
		Mari belajar "Golang".
	`
	fmt.Println(message)
	// petik backticks (`) u/ mmebuat yang di dalamnya tidak escape dan menjadi string semua

	var name = "John Doe"
	var age = "28"
	var sentence = `halo nama saya "` + name + `" dan berumur "` + age + `"`

	fmt.Println()
	fmt.Println(sentence)

	/* 5. Nil dan Zero Value */
	/*
		Zero Value :
			Nilai default yang ada di beberapa tipe data, nilai default yang dimaksud
			adalah ketika nilai dari suatu variable tidak ada. For Ex,

				- Zero value dari string adalah "" (string kosong).
				- Zero value dari bool adalah false.
				- Zero value dari tipe numerik non-desimal adalah 0.
				- Zero value dari tipe numerik desimal adalah 0.0.

		Nil Value :
			Nilai yang benar2 kosong jika tidak ada isinya (tidak mempunyai default value)
			nil tidak bisa digunakan pada tipe data diatas. Beberapa tipe data yang bisa
			pake nil,
				- pointer
				- tipe data fungsi
				- slice
				- map
				- channel
				- interface kosong atau interface{}
	*/

	/* 6. Konversi Data dengan Teknik Casting */

	/*
		Keyword tipe data bisa digunakan untuk casting, atau konversi antar tipe data.
		Cara penggunaannya adalah dengan menuliskan tipe data tujuan casting sebagai fungsi,
		lalu menyisipkan data yang akan dikonversi sebagai parameter fungsi tersebut.
	*/

	var a float64 = float64(24) //float64() => fungsi
	fmt.Println(a)              // 24

	var b int32 = int32(24.00)
	fmt.Println(b) // 24

	var str = "Halo"
	var c string = string(str[0])
	fmt.Println(c) // H

	/* 7. Packages strings dan strconv */

	/* Dari GO sudah ada package string yang terdapat beberapa fungsi u/ mengolah string */

	/* A. String Index */
	// Digunakan untuk mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama).

	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2
	// String "ha" berada pada posisi ke 2 dalam string "ethan hunt" (indeks dimulai dari 0). Jika diketemukan dua substring, maka yang diambil adalah yang pertama, contoh:

	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2) // 4
	// String "n" berada pada indeks 4 dan 8. Yang dikembalikan adalah yang paling kiri (paling kecil), yaitu 4.

	/* B. String Replace */
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	/* C. String Repeat */
	var halo = strings.Repeat("halo", 4)
	fmt.Println(halo) // halohalohalohalo

	/* D. String ToLower */
	var lower = strings.ToLower("aNJay")
	fmt.Println(lower) // anjay

	/* E. String ToUpper */
	var upper = strings.ToUpper("haha")
	fmt.Println(upper) // HAHA

	/* More Information about package string in https://pkg.go.dev/strings */

	/* Package Strconv */
	
}
