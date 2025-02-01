package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"main/utils"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appconfigdata"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	utils.Check(err)

	// Flags
	port := flag.Int("port", 8080, "")
	flag.Parse()

	// Initiate App Config
	startConfig()

	// Register the handlers
	http.HandleFunc("/", ok)
	http.HandleFunc("/config", configResource)

	// Start the HTTP server
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Staring server on port %v", *port)
	http.ListenAndServe(addr, nil)
}

func ok(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK\n")
}

var client *appconfigdata.Client

// var exit = make(chan bool)
var token *string

// var nextPollIntervalInSeconds int
var configStr string

func startConfig() {
	// Use this token only the first time
	initialConfigurationToken := startSession()
	token = initialConfigurationToken
	go getLatestConfigInfiniteLoop()
}

func startSession() *string {

	params := utils.GetParameters()

	cfg, err := config.LoadDefaultConfig(context.TODO())
	utils.Check(err)

	var minInterval int32 = 15

	client = appconfigdata.NewFromConfig(cfg)
	input := appconfigdata.StartConfigurationSessionInput{
		ApplicationIdentifier:                &params.AppId,
		ConfigurationProfileIdentifier:       &params.ConfigProfileId,
		EnvironmentIdentifier:                &params.EnvId,
		RequiredMinimumPollIntervalInSeconds: &minInterval,
	}
	output, err := client.StartConfigurationSession(context.TODO(), &input)
	utils.Check(err)

	return output.InitialConfigurationToken
}

func getLatestConfigInfiniteLoop() {
	for {
		fmt.Printf("Retrieving latest config...\n")
		input := appconfigdata.GetLatestConfigurationInput{
			ConfigurationToken: token,
		}
		output, err := client.GetLatestConfiguration(context.TODO(), &input)
		utils.Check(err)

		token = output.NextPollConfigurationToken
		latest := output.Configuration
		if len(latest) != 0 {
			configStr = string(latest)
			fmt.Printf("New latest configuration retrieved! %s\n", configStr)
			fmt.Println("New latest configuration retrieved!")
		} else {
			fmt.Println("Nothing changed, already using the latest configuration!")
		}

		interval := output.NextPollIntervalInSeconds
		duration := time.Second * time.Duration(interval)
		fmt.Printf("Sleeping... %s\n", duration)
		time.Sleep(duration)
	}
}

func configResource(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK\n")
}
