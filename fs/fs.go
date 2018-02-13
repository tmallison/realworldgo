package main

import (
	"flag"
	"net/http"
	"os"
	"fmt"
)

func main() {
	var dir string

	port := flag.String("port", "3000", "port to serve HTTP on")
	path := flag.String("path", "", "path to server")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}

	fmt.Println(dir)

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}