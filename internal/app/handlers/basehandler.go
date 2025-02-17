package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Twix298/shortener/internal/app/storage"
)

type BaseHandler struct {
	BaseUrl string
	Storage storage.Storage
}

func MakeBaseHandler(url string, strorage storage.Storage) *BaseHandler {
	return &BaseHandler{BaseUrl: url, Storage: strorage}
}

func (base *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Printf("POST Method\n")
		base.getShortUrl(w, r)
	} else if r.Method == http.MethodGet {
		fmt.Printf("Get Method\n")
		base.getFullUrl(w, r)
	} else {
		http.Error(w, "Отправлен неподходящий метод", http.StatusMethodNotAllowed)
	}
}

func (base *BaseHandler) getShortUrl(w http.ResponseWriter, r *http.Request) {
	bytes, _ := io.ReadAll(r.Body)
	stringUrl := string(bytes)
	fmt.Printf("%s\n", stringUrl)

	hash, err := base.Storage.SaveUrl(stringUrl)
	fmt.Printf("%s\n", hash)
	if err != nil {
		http.Error(w, "Ошибка создания url", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resp := base.BaseUrl + "/" + hash
	w.Write([]byte(resp))
}

func (base *BaseHandler) getFullUrl(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Path[1:]
	fmt.Printf("%s\n", hash)
	var err error
	resp, err := base.Storage.SearchUrl(hash)
	fmt.Printf("Responce searchUrl = %s\n", resp)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, "Ошибка декодирования", http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", resp)
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(resp))
}
