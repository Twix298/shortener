package main

import (
	"fmt"
	"net/http"

	"github.com/Twix298/shortener/internal/app/handlers"
	"github.com/Twix298/shortener/internal/app/storage"
	"github.com/Twix298/shortener/internal/app/url_generator"
)

func run() error {
	generatorUrl := url_generator.Generator{Lenght: 8}
	storage := storage.Instance(generatorUrl)
	host := "http://localhost:8080"
	baseHandler := handlers.BaseHandler{
		BaseUrl: host,
		Storage: storage,
	}
	return http.ListenAndServe(":8080", newRouter(&baseHandler))

}

func main() {
	fmt.Println("ok")
	if err := run(); err != nil {
		panic("unexpected error: " + err.Error())
	}

}
