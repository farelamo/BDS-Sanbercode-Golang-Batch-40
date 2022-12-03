package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

/*********** Soal 1 **********/
func deferSoal1(kalimat string, tahun int) {
	fmt.Println(kalimat, tahun)
}

func soal1(kalimat string, tahun int) {
	defer deferSoal1(kalimat, tahun)
	fmt.Println("========= Soal Nomor 1 ========")
}

/*********** Soal 2 **********/
func kelilingSegitigaSamaSisi(value int, condition bool) (string, error) {
	if condition {
		if value == 0 {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}

		return "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(value) + " cm adalah " + strconv.Itoa(value*3) + " cm", nil
	} else {
		if value > 0 {
			return strconv.Itoa(value), nil
		}

		defer handleError()
		panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}
}

func handleError() {
	message := recover()
	fmt.Println("Oops,", message)
}

/*********** Soal 3 **********/
func tambahAngka(angka int, angkaNow *int){
	*angkaNow += angka
}

func cetakAngka(angkaNow *int){
	fmt.Println("Angka sekarang :", *angkaNow)
}

/*********** Soal 4 **********/
func printInterval(phones *[]string){
	*phones = append(*phones, "Xiaomi")
	*phones = append(*phones, "Asus")
	*phones = append(*phones, "IPhone")
	*phones = append(*phones, "Samsung")
	*phones = append(*phones, "Oppo")
	*phones = append(*phones, "Realme")
	*phones = append(*phones, "Vivo")

	for i, phone := range *phones {
		time.Sleep(time.Second * 1)
		fmt.Println(i+1, phone)
	}
}

func main() {
	/* Untuk Soal Nomor 3 */
		// deklarasi variabel angka ini simpan di baris pertama func main
		angka := 1
	/**********************/

	// Soal 1
	soal1("Golang Backend Development", 2021)

	fmt.Println() // u/ jarak antar soal

	// Soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	fmt.Println() // u/ jarak antar soal

	// Soal 3
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	fmt.Println() // u/ jarak antar soal

	// Soal 4
	var phones = []string{}
	printInterval(&phones)
}
