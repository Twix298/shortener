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

func (base *BaseHandler) GetShortUrl(w http.ResponseWriter, r *http.Request) {
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

func (base *BaseHandler) GetFullUrl(w http.ResponseWriter, r *http.Request) {
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
