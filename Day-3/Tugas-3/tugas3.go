package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Soal 1
	var panjangPersegiPanjang string 	= "8"
	var lebarPersegiPanjang string 		= "5"

	var alasSegitiga string 			= "6"
	var tinggiSegitiga string 			= "7"

	pPersegiPanjang, _ 	:= strconv.Atoi(panjangPersegiPanjang)
	lPersegiPanjang, _ 	:= strconv.Atoi(lebarPersegiPanjang)
	aSegitiga, _ 		:= strconv.Atoi(alasSegitiga)
	tSegitiga, _ 		:= strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	luasPersegiPanjang      = pPersegiPanjang * lPersegiPanjang
	kelilingPersegiPanjang  = (pPersegiPanjang + lPersegiPanjang)/2
	luasSegitiga			= (aSegitiga * tSegitiga)/2

	fmt.Println(luasPersegiPanjang)
	fmt.Println(kelilingPersegiPanjang)
	fmt.Println(luasSegitiga)

	// Soal 2
	if nilaiJohn := 80; nilaiJohn >= 80 {
		fmt.Println("Nilai Jhon A")
	}else if(nilaiJohn < 80 && nilaiJohn >= 70){
		fmt.Println("Nilai Jhon B")
	}else if(nilaiJohn >= 60 && nilaiJohn < 70){
		fmt.Println("Nilai Jhon C")
	}else if(nilaiJohn >= 50 && nilaiJohn < 60) {
		fmt.Println("Nilai Jhon D")
	}else {
		fmt.Println("Nilai Jhon E")
	}

	if nilaiDoe := 50; nilaiDoe >= 80 {
		fmt.Println("Nilai Doe A")
	}else if(nilaiDoe < 80 && nilaiDoe >= 70){
		fmt.Println("Nilai Doe B")
	}else if(nilaiDoe >= 60 && nilaiDoe < 70){
		fmt.Println("Nilai Doe C")
	}else if(nilaiDoe >= 50 && nilaiDoe < 60) {
		fmt.Println("Nilai Doe D")
	}else {
		fmt.Println("Nilai Doe E")
	}

	// Soal 3
	var tanggal = 9
	var bulan   = 11
	var tahun 	= 2001

	switch bulan {
		case 1:  fmt.Println(tanggal, "Januari", tahun)
		case 2:  fmt.Println(tanggal, "Februari", tahun)
		case 3:  fmt.Println(tanggal, "Maret", 	tahun)
		case 4:  fmt.Println(tanggal, "April", tahun)
		case 5:  fmt.Println(tanggal, "Mei", tahun)
		case 6:  fmt.Println(tanggal, "Juni", tahun)
		case 7:  fmt.Println(tanggal, "Juli", tahun)
		case 8:  fmt.Println(tanggal, "Agustus", tahun)
		case 9:  fmt.Println(tanggal, "September", tahun)
		case 10: fmt.Println(tanggal, "Oktober", tahun)
		case 11: fmt.Println(tanggal, "November", tahun)
		case 12: fmt.Println(tanggal, "Desember", tahun)
	}

	// Soal 4
	if tahunLahir := 2001; (tahunLahir >= 1944 && tahunLahir <= 1964) {
		fmt.Println("Baby boomer, kelahiran 1944 s.d 1964")
	}else if(tahunLahir >= 1965 && tahunLahir <= 1979){
		fmt.Println("Generasi X, kelahiran 1965 s.d 1979")
	}else if(tahunLahir >= 1980 && tahunLahir <= 1994){
		fmt.Println("Generasi Y (Millenials), kelahiran 1980 s.d 1994")
	}else {
		fmt.Println("Generasi Z, kelahiran 1995 s.d 2015")
	}
}
