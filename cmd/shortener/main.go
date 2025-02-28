package main

import (
	"fmt"
	"net/http"

	"github.com/Twix298/shortener/internal/app/config"
	"github.com/Twix298/shortener/internal/app/handlers"
	"github.com/Twix298/shortener/internal/app/storage"
	"github.com/Twix298/shortener/internal/app/url_generator"
)

func run() error {
	generatorUrl := url_generator.Generator{Lenght: 8}
	storage := storage.Instance(generatorUrl)
	baseHandler := handlers.BaseHandler{
		BaseUrl: config.BaseURL,
		Storage: storage,
	}
	return http.ListenAndServe(config.Port, newRouter(&baseHandler))

}

func main() {
	fmt.Println("ok")
	config.Parse()
	if err := run(); err != nil {
		panic("unexpected error: " + err.Error())
	}

}
