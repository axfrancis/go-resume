package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go")
}

func resume(w http.ResponseWriter, r *http.Request) {
	pdf, err := generatePDF()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdf)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/resume", resume)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
