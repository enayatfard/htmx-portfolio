package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Hello, htmx!</h1>")
    })

    corsHandler := func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "*")
            w.Header().Set("Access-Control-Allow-Headers", "*")

            next.ServeHTTP(w, r)
        })
    }

    http.ListenAndServe(":8080", corsHandler(http.DefaultServeMux))
}
