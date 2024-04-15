package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	message := os.Getenv("HELLO_MESSAGE!")
	if message == "" {
		message = "Hello, World! (no set env)"
	}
	fmt.Fprint(w, message)
}

func main() {
	http.HandleFunc("/", helloWorld)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running at http://localhost:%s\n"), port)
	err := http.ListenAndServe(":"+port, nil)
	if port =="" {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
