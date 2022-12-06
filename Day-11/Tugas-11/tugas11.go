package main

import (
	"fmt"
	"sync"
	"sort"
  )

func printText(i int, text string, wp *sync.WaitGroup){
	fmt.Println(i+1, ".", text)
	wp.Done()
}

func getMovies(moviesChannel chan string, movies ...string){
	for _, movie := range movies {
		moviesChannel <- movie
	}
	close(moviesChannel)
}

func main(){
	
	// Soal 1
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	var wp sync.WaitGroup

	sort.Strings(phones)
	for i, phone := range phones {
		wp.Add(1)
		go printText(i, phone, &wp)
		wp.Wait()
	}

	fmt.Println() // untuk jarak soal

	// Soal 2
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
}

