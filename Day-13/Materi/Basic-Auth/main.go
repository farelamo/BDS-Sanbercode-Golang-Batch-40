package main

import (
	"fmt"
	"net/http"
)

// Auth Middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau password tidak bole kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}

		w.Write([]byte("Username atau password salah"))
		return
	})
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("<h1>Anda berhasil mengakses fungsi getMovie()</h1>"))
	}
}


func main() {
	// konfigurasi server
	server := &http.Server {
		Addr: ":8000",
	}

	// Routing
	http.Handle("/", Auth(http.HandlerFunc(getMovie)))

	// Run server
	fmt.Println("Server running at http://localhost:8000")
	server.ListenAndServe()
}