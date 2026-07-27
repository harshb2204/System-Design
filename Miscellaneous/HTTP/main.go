package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from server"))
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Listening on port 1729")
	err := http.ListenAndServe(":1729", nil)

	if err != nil {
		panic(err)
	}

}
