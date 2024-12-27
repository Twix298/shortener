package main

import (
	"fmt"
	"io"
	"net/http"
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
	resp := "ok"
	w.Header().Set("content-type", "application/json")
	// устанавливаем код 200
	w.WriteHeader(http.StatusCreated)
	// пишем тело ответа
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
