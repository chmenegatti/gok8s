package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if name == "" {
		name = "John Doe"
	}
	if !ok {
		fmt.Println("Erro")
	}

	w.Write([]byte(fmt.Sprintf("<h1>Hello, %s</h1>", name)))
}
