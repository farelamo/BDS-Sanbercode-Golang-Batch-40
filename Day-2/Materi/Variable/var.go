package main

import ("fmt")

func main (){
	/* Contoh 1 */
	var text = "Hello World"
	fmt.Println(text)

	text = "Hello Farelamo"
	fmt.Println(text)

	/* Contoh 2 (Dengan tipe data) */
	var text1 string
	text1 = "Hello World"
	fmt.Println(text1)

	var text2 string = "ini text 2"
	text2 = "text 2 dirubah"
	fmt.Println(text2)

	/* Contoh 3 (Dengan :=) */
	text3 := "ini text 3"
	text3 = "ini text 3 dirubah"
	fmt.Println(text3)
}