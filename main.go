package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Haha")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
