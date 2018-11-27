package handler

import (
	"net/http"
	"fmt"
)

func HelthCheck(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "Hello!")
	return nil
}
