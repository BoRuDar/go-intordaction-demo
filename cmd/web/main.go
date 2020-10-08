package main

import (
	"fmt"
	"net/http"

	"github.com/BoRuDar/go-introduction-demo/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	services.NewGitHub(http.DefaultClient)

	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", helloHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	_ = srv.ListenAndServe()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintf(w, "Hello: %v\n", vars["name"])
}
