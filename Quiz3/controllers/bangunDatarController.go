package controllers

import (
	"Quiz3/library"
	"encoding/json"
	"net/http"
	"strconv"
)

var BangunDatar library.HitungBangunDatar

func SegitigaSamaSisi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if r.Header.Get("Content-Type") != "application/json" {
			alas, _		:= strconv.Atoi(r.PostFormValue("alas"))
			tinggi, _ 	:= strconv.Atoi(r.PostFormValue("tinggi"))

			if r.PostFormValue("hitung") == "luas" {
				r := library.SegitigaSamaSisi{
						Alas: alas, 
						Tinggi: tinggi,
						"status" : true,
						"message": "Successfully Calculated Data",
						"Result" : BangunDatar.Luas(),
				}
				
				data, _ := json.Marshal(r)
				w.Write(data)
				return
			} else {
				r := library.SegitigaSamaSisi{
					Alas: alas, 
					Tinggi: tinggi,
					"status" : true,
					"message": "Successfully Calculated Data",
					"Result" : BangunDatar.keliling(),
				}
				data, _ := json.Marshal(r)
				w.Write(data)
				return
			}
		}
	}else {
		http.Error(w, "Method Salah", http.StatusMethodNotAllowed)
		return
	}
}

func SegitigaSamaSisi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if r.Header.Get("Content-Type") != "application/json" {
			alas, _		:= strconv.Atoi(r.PostFormValue("alas"))
			tinggi, _ 	:= strconv.Atoi(r.PostFormValue("tinggi"))

			if r.PostFormValue("hitung") == "luas" {
				r := library.SegitigaSamaSisi{
						Alas: alas, 
						Tinggi: tinggi,
						"status" : true,
						"message": "Successfully Calculated Data",
						"Result" : BangunDatar.Luas(),
				}
				
				data, _ := json.Marshal(r)
				w.Write(data)
				return
			} else {
				r := library.SegitigaSamaSisi{
					Alas: alas, 
					Tinggi: tinggi,
					"status" : true,
					"message": "Successfully Calculated Data",
					"Result" : BangunDatar.keliling(),
				}
				data, _ := json.Marshal(r)
				w.Write(data)
				return
			}
		}
	}else {
		http.Error(w, "Method Salah", http.StatusMethodNotAllowed)
		return
	}
}

func SegitigaSamaSisi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if r.Header.Get("Content-Type") != "application/json" {
			alas, _		:= strconv.Atoi(r.PostFormValue("alas"))
			tinggi, _ 	:= strconv.Atoi(r.PostFormValue("tinggi"))

			if r.PostFormValue("hitung") == "luas" {
				r := library.SegitigaSamaSisi{
						Alas: alas, 
						Tinggi: tinggi,
						"status" : true,
						"message": "Successfully Calculated Data",
						"Result" : BangunDatar.Luas(),
				}
				
				data, _ := json.Marshal(r)
				w.Write(data)
				return
			} else {
				r := library.SegitigaSamaSisi{
					Alas: alas, 
					Tinggi: tinggi,
					"status" : true,
					"message": "Successfully Calculated Data",
					"Result" : BangunDatar.keliling(),
				}
				data, _ := json.Marshal(r)
				w.Write(data)
				return
			}
		}
	}else {
		http.Error(w, "Method Salah", http.StatusMethodNotAllowed)
		return
	}
}

func SegitigaSamaSisi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if r.Header.Get("Content-Type") != "application/json" {
			alas, _		:= strconv.Atoi(r.PostFormValue("alas"))
			tinggi, _ 	:= strconv.Atoi(r.PostFormValue("tinggi"))

			if r.PostFormValue("hitung") == "luas" {
				r := library.SegitigaSamaSisi{
						Alas: alas, 
						Tinggi: tinggi,
						"status" : true,
						"message": "Successfully Calculated Data",
						"Result" : BangunDatar.Luas(),
				}
				
				data, _ := json.Marshal(r)
				w.Write(data)
				return
			} else {
				r := library.SegitigaSamaSisi{
					Alas: alas, 
					Tinggi: tinggi,
					"status" : true,
					"message": "Successfully Calculated Data",
					"Result" : BangunDatar.keliling(),
				}
				data, _ := json.Marshal(r)
				w.Write(data)
				return
			}
		}
	}else {
		http.Error(w, "Method Salah", http.StatusMethodNotAllowed)
		return
	}
}

func Lingkaran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if r.Header.Get("Content-Type") != "application/json" {
			jari, _		:= strconv.Atoi(r.PostFormValue("jari"))

			if r.PostFormValue("hitung") == "luas" {
				r := library.Lingkaran{
						Alas: jari,
						"status" : true,
						"message": "Successfully Calculated Data",
						"Result" : BangunDatar.Luas(),
				}
				
				data, _ := json.Marshal(r)
				w.Write(data)
				return
			} else {
				r := library.Lingkaran{
					JariJari: alas,
					"status" : true,
					"message": "Successfully Calculated Data",
					"Result" : BangunDatar.keliling(),
				}
				data, _ := json.Marshal(r)
				w.Write(data)
				return
			}
		}
	}else {
		http.Error(w, "Method Salah", http.StatusMethodNotAllowed)
		return
	}
}