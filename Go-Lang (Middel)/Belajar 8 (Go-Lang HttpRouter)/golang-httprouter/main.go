package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ================================================== Router ==================================================
func main() {
	router := httprouter.New() //turunan atau implementasi handler

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello Ini Adalah GET")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
