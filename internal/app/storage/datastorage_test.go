package storage

import (
	"fmt"
	"testing"

	"github.com/Twix298/shortener/internal/app/url_generator"
)

func TestStorage_SaveUrl(t *testing.T) {
	type fields struct {
		Url       map[string]string
		Generator url_generator.Generator
	}
	type args struct {
		url string
	}

	generatorUrl := url_generator.Generator{Lenght: 8}
	urlStorage := make(map[string]string)
	fueld := fields{
		urlStorage,
		generatorUrl,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test_1",
			fields:  fueld,
			args:    args{"www.yandex.ru"},
			want:    "d3d3Lnlh",
			wantErr: false,
		},
		{
			name:    "test_2",
			fields:  fueld,
			args:    args{"www.google.com"},
			want:    "d3d3Lmdv",
			wantErr: false,
		},
		{
			name:    "test_3",
			fields:  fueld,
			args:    args{"www.google.com"},
			want:    "d3d3Lmdv",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Storage{
				Url:       tt.fields.Url,
				Generator: tt.fields.Generator,
			}
			got, err := d.SaveUrl(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.SaveUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.SaveUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_SearchUrl(t *testing.T) {
	type fields struct {
		Url       map[string]string
		Generator url_generator.Generator
	}
	type args struct {
		hash string
	}

	generatorUrl := url_generator.Generator{Lenght: 8}
	urlStorage := make(map[string]string)

	sliceUrl := []string{"www.yandex.ru", "www.google.com"}

	for _, val := range sliceUrl {
		urlStorage[generatorUrl.EncodeURL(val)] = val
	}
	for k, v := range urlStorage {
		fmt.Printf("k = %s, v = %s\n", k, v)
	}

	fueld := fields{
		urlStorage,
		generatorUrl,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test_1",
			fields:  fueld,
			args:    args{"d3d3Lnlh"},
			want:    "www.yandex.ru",
			wantErr: false,
		},
		{
			name:    "test_2",
			fields:  fueld,
			args:    args{"d3d3Lmdv"},
			want:    "www.google.com",
			wantErr: false,
		},
		{
			name:    "test_3",
			fields:  fueld,
			args:    args{"d3d3L"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Storage{
				Url:       tt.fields.Url,
				Generator: tt.fields.Generator,
			}
			got, err := d.SearchUrl(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.SearchUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.SearchUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
