package config

import (
	"flag"
	"os"
)

var (
	Port    = "localhost:8080"
	BaseURL = "http://localhost:8080/"
)

func Parse() {
	flag.StringVar(&Port, "a", Port, "port to run server")
	flag.StringVar(&BaseURL, "b", BaseURL, "base URL for shorten URL response")

	flag.Parse()

	if val := os.Getenv("SERVER_ADDRESS"); val != "" {
		Port = val
	}
	if val := os.Getenv("BASE_URL"); val != "" {
		BaseURL = val
	}
}
