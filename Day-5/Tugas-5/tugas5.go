package main

import (
	"fmt"
)

func main(){

	// Soal 1
	panjang := 12
	lebar 	:= 4
	tinggi  := 8

	luas 	  := luasPersegiPanjang(panjang, lebar)
	keliling  := kelilingPersegiPanjang(panjang, lebar)
	volume 	  := volumePersegiPanjang(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)


	// Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "papaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// Soal 4
	var dataFilm = []map[string]string{}
	var tambahDataFilm = func(title, jam, genre, tahun string){
		dataFilm = append(dataFilm, map[string]string{"gender": genre, "jam": jam, "tahun": tahun, "title": title})
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")

	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

func luasPersegiPanjang(p int, l int) int {
	luas := p * l
	return luas
}

func kelilingPersegiPanjang(p int, l int) int {
	keliling := (p + l)/2
	return keliling
}

func volumePersegiPanjang(p int, l int, t int) int {
	volume := p * l * t
	return volume
}

func introduce(name, gender, job, age string) (result string) {
	if gender == "laki-laki" {
		result = "Pak " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
		return
	}else {
		return "Bu " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	}
}

func buahFavorit(name string, fruits ...string) (word string) {
	word = "halo nama saya " + name + " dan buah favorit saya adalah "
	for _, fruit := range fruits {
		word += "'" + fruit + "', "
	}
	return
}