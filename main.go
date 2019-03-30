package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
)

var Version string

func main() {

	fmt.Println("version:", Version)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, hash(r.URL.Query().Get("q")))
	})

	http.ListenAndServe(":8080", nil)
}

func hash(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}
