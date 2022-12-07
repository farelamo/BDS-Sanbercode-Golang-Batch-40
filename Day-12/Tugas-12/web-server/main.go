package webserver

import (
	"fmt"
	"math"
	"net/http"
)

func volume(r float64, t float64) float64 {
	return math.Phi * math.Pow(r, 2) * t
}

func luasAlas(r float64) float64 {
	return math.Phi * math.Pow(r, 2)
}

func kelilingAlas(r float64) float64 {
	return 2 * math.Phi * r
}

func main() {
	var jari = 7.0
	var t = 10.0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "jariJari :", jari, ",", "tinggi :", t, "volume :", volume(jari, t), "luas :", luasAlas(jari), "keliling :", kelilingAlas(jari))
	})

	fmt.Println("starting web server at http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
