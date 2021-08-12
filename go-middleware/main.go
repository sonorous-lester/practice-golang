package main

import (
	"fmt"
	"go-middleware/middleware"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/home", home)

	// only one middleware to act on this request
	mux.Handle("/private", goMiddleware(http.HandlerFunc(secret)))

	// two middlewares to act on every http request
	err := http.ListenAndServe(":4000", middleware.LogRequestMiddleware(middleware.SecureHeadersMiddleware(mux)))

	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go Middleware"))
}

func secret(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("No one will know this secret.")
}

func goMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("Hello, I'm go middleware.")
		next.ServeHTTP(writer, request)
	})
}
