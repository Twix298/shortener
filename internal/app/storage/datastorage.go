package storage

import (
	"errors"

	"github.com/Twix298/shortener/internal/app/url_generator"
)

type Storage struct {
	Url       map[string]string
	Generator url_generator.Generator
}

func Instance(gen url_generator.Generator) Storage {
	return Storage{
		Url:       make(map[string]string),
		Generator: gen}
}

func (d *Storage) SaveUrl(url string) (string, error) {
	hash := d.Generator.EncodeURL(url)
	d.Url[hash] = url
	return hash, nil
}

func (d *Storage) SearchUrl(hash string) (string, error) {
	val, ok := d.Url[hash]
	if ok {
		return val, nil
	}
	return "", errors.New("value not exist")
}
