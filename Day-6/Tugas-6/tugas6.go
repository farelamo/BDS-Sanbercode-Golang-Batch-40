package main

import (
	"fmt"
)

func main() {
	// Soal 1

	// var luasLigkaran float64
	// var kelilingLingkaran float64

	// Soal 2

	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3

	// var buah = []string{}
	// pushBuah(&buah)

	// Soal 4
	// var dataFilm = []map[string]string{}

	// var tambahDataFilm = func(title, jam, genre, tahun string, dataFilm ...*string){
	// 	*dataFilm = append(*dataFilm, map[string]string{"gender": genre, "jam": jam, "tahun": tahun, "title": title})
	// }

	// tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	// tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	// tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	// tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// for _, item := range dataFilm {
	// 	fmt.Println(item)
	// }

}

func introduce(sentence *string, name, gender, job, age string) {
	if gender == "laki-laki" {
		*sentence = "Pak " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	} else {
		*sentence = "Bu " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	}
}

// func pushBuah() (buah ...string){

// }
