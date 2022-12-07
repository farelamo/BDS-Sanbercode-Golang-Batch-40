package main

import (
	"fmt"
	"net/http"
)

// Membuat function middleware
func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Ini dari middleware Log.... \n")
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Membuat middleware cek login
func cekLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("token") != "12345" {
			fmt.Fprintf(w, "Token tidak tersedia atau salah \n")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Function u/ get data
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Ini dari function get Movies()"))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Anda berhasil mengakses fungsi getBooks() </h1>"))
}

func main(){
	// konfigurasi server
	server := &http.Server{
		Addr: ":8000",
	}

	// Routing
	http.Handle("/", log(http.HandlerFunc(getMovies)))
	http.Handle("/books", cekLogin(http.HandlerFunc(getBooks)))

	// Run server
	fmt.Println("Server is running at http://localhost:8000")
	server.ListenAndServe()
}