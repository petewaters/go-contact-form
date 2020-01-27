package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		fmt.Printf("Name: %s\n", r.FormValue("name"))
		fmt.Printf("Email: %s\n", r.FormValue("email"))
		fmt.Printf("Message: %s\n", r.FormValue("message"))
	})

	http.ListenAndServe(":8000", mux)
}
