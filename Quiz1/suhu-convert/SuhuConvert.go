package main

import ("fmt")

func main(){
	cetak(10)
}

func cetak(suhu int){
	var convert = convertFarenheit(suhu)
	fmt.Println("Suhu:", convert)
}

func convertFarenheit(celcius int) int {
	return (9/5) * celcius + 32
}