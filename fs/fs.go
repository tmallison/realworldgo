package main

import (
	"net/http"
	"os"
	"fmt"
)

func main() {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
}