package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Println("starting to serve")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This page is served by instance number " + os.Args[1] + "."))
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
