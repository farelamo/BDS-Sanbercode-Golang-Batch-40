package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type NilaiMahasiswa struct {
	Nama        string	`json:"nama"`
	MataKuliah 	string	`json:"mata_kuliah"`
	IndeksNilai string	`json:"indeks_nilai"`
	Nilai 		uint	`json:"nilai"`
	ID 			uint	`json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

// Auth Middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password Kosong"))
			return 
		}

		if user == "admin" && pass == "admin" {
			next.ServeHTTP(w, r)
			return
		}

		w.Write([]byte("Username atau password salah"))
		return
	})
}

// GET
func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {

		encodeData, err := json.Marshal(&nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encodeData)
		return
	}
	http.Error(w, "Method Salah", http.StatusMethodNotAllowed)
}

// POST
func addNilaiMahasiswa(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var tmp NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON 	:= json.NewDecoder((r.Body))
			if err 	    := decodeJSON.Decode(&tmp); err != nil {
				log.Fatal(err)
			}
		}else {
			http.Error(w, "Data harus berupa JSON", http.StatusBadRequest)
			return
		}
		
		if tmp.Nilai >= 80 {
			tmp.IndeksNilai = "A"
		}else if(tmp.Nilai < 80 && tmp.Nilai >= 70){
			tmp.IndeksNilai = "B"
		}else if(tmp.Nilai < 70 && tmp.Nilai >= 60){
			tmp.IndeksNilai = "C"
		}else if(tmp.Nilai < 60 && tmp.Nilai >= 50){
			tmp.IndeksNilai = "D"
		}else {
			tmp.IndeksNilai = "E"
		}

		tmp.ID = uint(len(nilaiNilaiMahasiswa) + 1)

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, tmp)
		result, _ 			:= json.Marshal(nilaiNilaiMahasiswa)
		w.Write(result)
		return
	}

	http.Error(w, "Method Salah", http.StatusMethodNotAllowed)
}

func main(){
	// Route
	http.HandleFunc("/", getNilaiMahasiswa)
	http.Handle("/create", Auth(http.HandlerFunc(addNilaiMahasiswa)))

	fmt.Println("Server running at http://localhost:8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}