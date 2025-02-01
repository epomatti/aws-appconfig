package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := fmt.Sprintf(":%d", 8080)
	log.Printf("Staring server on port %v", 8080)
	http.ListenAndServe(addr, nil)
}
