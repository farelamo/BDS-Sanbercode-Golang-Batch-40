package main

import (
	"fmt"
	"sort"
)

type car struct{
	name, brand string
	year int
  }

func main(){
	phones := []string{"asus", "iPhone", "Samsung", "Xiaomi", "Realme", "motorola"}
	sort.Strings(phones)
	fmt.Println(phones)

	angka := 20

	var angkaBaru *int = &angka

	*angkaBaru += 5

	angka = 20 - *angkaBaru
	
	angka += 3

	*angkaBaru += 20

	fmt.Println(angka)
}