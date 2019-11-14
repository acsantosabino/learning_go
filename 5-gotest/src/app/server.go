package main

import (
	"fmt"
	"model/user"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	u := user.NewUser()
	u.SetName("Arthur")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, welcome(u.GetName()))
		// "Welcome to this life-changing API.\nIts the best API, its true, all other API's are fake.")
	})

	fmt.Println("Server listening!")
	http.ListenAndServe(":80", r)
}
