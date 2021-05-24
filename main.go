package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Error Ocurred", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/bye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Bye World")
	})

	http.ListenAndServe(":9090", nil)
}
