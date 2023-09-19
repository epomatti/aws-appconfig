package main

import (
	"fmt"
	"main/utils"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	utils.Check(err)

	port := utils.GetPort()

	http.HandleFunc("/hello", hello)

	// Server
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK\n")
}
