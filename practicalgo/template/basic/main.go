package main

import (
	"html/template"
	"log"
	"net/http"
)

type Product struct {
	Name        string
	Price       string
	Description string
}

func redTeaPotHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("/home/haltin/studies/practicalgo/views/product.html")

	if err != nil {
		http.Error(w, "Failed to parse the template", http.StatusInternalServerError)
		log.Printf("Failed to parse template: %v", err)
		return
	}
	teaPot := Product{Name: "Red Tea Pot 250ml", Description: " llll Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididun tut labore et dolore", Price: "19.99"}

	err = tmpl.Execute(w, teaPot)
	// handle error
	if err != nil {
		http.Error(w, "Failed to execute the template", http.StatusInternalServerError)
		log.Printf("Failed to execute template: %v", err)
	}
}

func main() {
	http.HandleFunc("/red-tea-pot", redTeaPotHandler)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}

}
