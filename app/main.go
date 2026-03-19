package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello this is ok\n")
    })

	port := ":8080"

	log.Println("Server running on port", port)
    log.Fatal(http.ListenAndServe(port, nil))
}