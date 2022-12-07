package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

type Movie struct {
	ID 		int 	`json:"id"`
	Title 	string	`json:"title"`
	Year    int		`json:"year"`
}

func Movies() []Movie {
	movies := []Movie{
		{1, "Spider-Man", 2022},
		{2, "John Wick", 2014},
		{3, "Avanger", 2018},
		{4, "Logan", 2017},
	}

	return movies
}

func getMovies(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		movies := Movies()
		dataMovies, err := json.Marshal(movies)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMovies)
		return
	}
	http.Error(w, "ERROR.....", http.StatusNotFound)
}

func postMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var Mov Movie
	if r.Method == "POST" {
		if r.Header.Get("Content-type") == "application/json" {
			// parse dari json
			decodeJSON 	:= json.NewDecoder(r.Body)
			if err 		:= decodeJSON.Decode(&Mov); err != nil {
				log.Fatal(err)
			}
		}else {
			// parse dari form
			getID 		:= r.PostFormValue("id")
			id, _ 	 	:= strconv.Atoi(getID)
			getTitle	:= r.PostFormValue("title")
			getYear		:= r.PostFormValue("year")
			year, _ 	:= strconv.Atoi(getYear)

			Mov 		= Movie{
				ID		: id,
				Title	: getTitle,
				Year	: year,
			}
		}

		dataMovie, _ := json.Marshal(Mov) //to byte
		w.Write(dataMovie)				  // cetak di browser
		return	
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

func main(){
	// Route
	http.HandleFunc("/movies", getMovies)
	http.HandleFunc("/create", postMovie)

	fmt.Println("Server running at http://localhost:8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

