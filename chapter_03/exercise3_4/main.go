package main

import (
	"log"
	"net/http"
	"exercise3_4/svg"
)

func main() {
	http.HandleFunc("/fig.svg", logRequest(handler))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(svg.Svg()))
}

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL)
		next(w, r)
	}
}