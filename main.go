package main

import (
	"log"
	"net/http"
)

func main() {
	// Sajikan folder statis
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Sajikan file index.html untuk root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	log.Println("Server berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
