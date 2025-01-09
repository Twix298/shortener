package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Twix298/shortener/internal/app/url_generator"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Отправлен неподходящий метод", http.StatusMethodNotAllowed)
	}
	fmt.Print(r.Header)
	fmt.Println('\n')
	bytes, _ := io.ReadAll(r.Body)
	fmt.Print(string(bytes))
	fmt.Println('\n')
	resp := url_generator.EncodeURL(string(bytes))
	fmt.Print(resp)
	fmt.Println('\n')
	w.Header().Set("content-type", "application/json")
	// устанавливаем код 200
	w.WriteHeader(http.StatusCreated)
	resp = "http://localhost:8080/" + resp
	w.Write([]byte(resp))
}

func main() {
	fmt.Println("ok")
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
