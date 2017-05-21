package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Secret            string
	Access            string
	SQSURL            string
	Region            string
	ConnectionTimeout time.Duration
	WaitTime          int64
	PostEndpoint      string
	PostHost          string
	MaxSleep          time.Duration
	ContentType       string
	Connections       int64
}

// default values
const (
	defaultContentType       = "application/json"
	defaultRegion            = "us-east-1"
	defaultSleep             = time.Duration(300) * time.Second
	defaultConnectionTimeout = time.Duration(300) * time.Second
	defaultConnections       = 1
	defaultWaitTime          = 20
)

// Convert an Env var into an int64 or thrown error
func getEnvVarAsNum(envVar string) (int64, error) {
	val := os.Getenv(envVar)
	if val == "" {
		return 0, fmt.Errorf("The ENV variable, %s, was not found", envVar)
	}
	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func initConfig() {
	cfg.Access = os.Getenv("ACCESS")
	cfg.Secret = os.Getenv("SECRET")

	cfg.SQSURL = os.Getenv("SQS_URL")
	if cfg.SQSURL == "" {
		fmt.Println("Error SQS_URL not found.")
		os.Exit(1)
	}

	cfg.PostEndpoint = os.Getenv("POST_ENDPOINT")
	if cfg.PostEndpoint == "" {
		fmt.Println("Error: POST_ENDPOINT not found.")
		os.Exit(1)
	}

	cfg.PostHost = os.Getenv("POST_HOST")
	if cfg.PostHost == "" {
		cfg.PostHost = "http://127.0.0.1:80"
	}

	cfg.Region = os.Getenv("REGION")
	if cfg.Region == "" {
		cfg.Region = defaultRegion
	}

	cfg.ContentType = os.Getenv("CONTENT_TYPE")
	if cfg.ContentType == "" {
		cfg.ContentType = defaultContentType
	}

	waitTime, err := getEnvVarAsNum("WAIT_TIME")
	if err != nil {
		cfg.WaitTime = defaultWaitTime
	} else {
		cfg.WaitTime = waitTime
	}

	sleep, err := getEnvVarAsNum("MAX_SLEEP")
	if err != nil {
		cfg.MaxSleep = defaultSleep
	} else {
		cfg.MaxSleep = time.Duration(sleep) * time.Second
	}

	connectionTimeout, err := getEnvVarAsNum("CONNECTION_TIMEOUT")
	if err != nil {
		cfg.ConnectionTimeout = defaultConnectionTimeout
	} else {
		cfg.ConnectionTimeout = time.Duration(connectionTimeout) * time.Second
	}

	connections, err := getEnvVarAsNum("CONNECTIONS")
	if err != nil {
		cfg.Connections = defaultConnections
	} else {
		cfg.Connections = connections
	}
}
