package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
