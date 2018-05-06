package gzip_test

import (
	"fmt"
	"net/http"

	"github.com/zengabor/gzip"
)

func Example() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})

	http.ListenAndServe(":8080", gzip.Handle(mux))
}

func ExampleHandle() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})

	http.ListenAndServe(":8080", gzip.Handle(mux))
}

func ExampleHandleFunc() {
	http.Handle("/", gzip.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	}))

	http.ListenAndServe(":8080", nil)
}
