package main

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})
	http.ListenAndServe(":8091", nil)
}


