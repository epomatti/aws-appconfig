package main

import (
	"context"
	"fmt"
	"main/utils"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	utils.Check(err)

	port := utils.GetPort()

	http.HandleFunc("/", ok)
	http.HandleFunc("/config", configFunc)

	// Server
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}

func ok(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK\n")
}

func configFunc(w http.ResponseWriter, r *http.Request) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	utils.Check(err)

	client := appconfig.NewFromConfig(cfg)
	println(client)

	fmt.Fprint(w, "OK\n")
}
