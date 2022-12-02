package main

import (
	"Materi/library"
	. "Materi/library" 
	// u/ menjadika package main dan library selevel, 
	// jadi tdk perlu panggil nama package nya dulu kalo mau pake

	f "fmt" // menggunakan alias u/ penyingkatan pemanggilan
)

func main() {
	library.SayHello("John")
	// library.introduce("John")
	
	SayHello("keke") // tdk perlu panggil package dulu

	f.Println("Coba Alias") // pemanggilan singkat dengan alias

	//sayHai("Riana") 
	// u/ akses unexported pada 1 package yang sama (main) pada 1 folder 
	// (cara run ada di go build atau go run *.go)

	
}