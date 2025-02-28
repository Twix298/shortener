package config

import (
	"flag"
)

var (
	Port    = "localhost:8080"
	BaseURL = "http://localhost:8080/"
)

func Parse() {
	flag.StringVar(&Port, "a", Port, "port to run server")
	flag.StringVar(&BaseURL, "b", BaseURL, "base URL for shorten URL response")

	flag.Parse()
}
