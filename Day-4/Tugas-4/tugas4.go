package main

import (
	"fmt"
)

func main(){

	// Soal 1
	for i := 1; i <= 20; i++ {
		
		if i % 2 != 0 && i % 3 == 0 {
			fmt.Println(i, "- I Love Coding")
		}else if i % 2 != 0 {
			fmt.Println(i, "- Santai")
		} else {
			fmt.Println(i, "- Berkualitas")
		}
	}

	// Soal 2
	for i := 7; i > 0; i-- {
		for j := 7; j >= i; j-- {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// Soal 3
	var kalimat 	= [...] string {"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var newKalimat 	= [...] string {kalimat[2], kalimat[3], kalimat[4], kalimat[5], kalimat[6]}
	fmt.Println(newKalimat) // [saya sangat senang belajar golang]

	// Soal 4
	var sayuran	= []string{}
	sayuran		= append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
	
	for i, sayur := range sayuran {
		fmt.Println(i+1, ".", sayur)
	}

	// Soal 5
	var satuan = map[string]int {
		"panjang"	: 7,
		"lebar"		: 4,
		"tinggi"	: 6,
	}

	for _, data := range satuan {
		if satuan["panjang"] == data {
			fmt.Println("panjang =", data)
		} else if satuan["lebar"] == data {
			fmt.Println("lebar =", data)
		} else if satuan["tinggi"] == data {
			fmt.Println("tinggi =", data)
		}
	}
	fmt.Println("Volume Balok =", satuan["panjang"] * satuan["lebar"] * satuan["tinggi"])
}