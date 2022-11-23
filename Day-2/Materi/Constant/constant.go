package main

import (
	"fmt"
)

func main() {
	const judul = "ini judul"
	fmt.Println(judul)
	// Jika judul dibawah ini di uncomment bakal error karna judul itu constant (nilai tetap)
	// judul = "haloo"
}