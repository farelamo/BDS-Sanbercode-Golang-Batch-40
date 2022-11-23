package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	// Soal 1
	text1 := "Bootcamp"
	text2 := "Digital"
	text3 := "Skill"
	text4 := "Sanbercode"
	text5 := "Golang"

	fmt.Println(text1, text2, text3, text4, text5)

	// Soal 2
	halo := "Halo Dunia"
	resultReplace := strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(resultReplace)

	//  Soal 3
	var kataPertama = "saya";
	var kataKedua = "senang";
	var kataKetiga = "belajar";
	var kataKeempat = "golang";

	kataKedua  	= strings.Replace(kataKedua, "s", "S", 1)
	kataKetiga 	= strings.Replace(kataKetiga, "r", "R", 1)
	kataKeempat = strings.ToUpper(kataKeempat)

	fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)

	// Soal 4
	var angkaPertama= "8";
	var angkaKedua= "5";
	var angkaKetiga= "6";
	var angkaKeempat = "7";

	num1, _ := strconv.Atoi(angkaPertama)
	num2, _ := strconv.Atoi(angkaKedua)
	num3, _ := strconv.Atoi(angkaKetiga)
	num4, _ := strconv.Atoi(angkaKeempat)

	fmt.Println(num1 + num2 + num3 + num4)

	// Soal 5
	kalimat := "halo halo bandung"
	angka 	:= 2021

	kalimat = strings.Replace(kalimat, "halo", "Hi", 2)
	fmt.Println(kalimat, "-", angka)
}