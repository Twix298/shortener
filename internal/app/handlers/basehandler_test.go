package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Twix298/shortener/internal/app/storage"
	"github.com/Twix298/shortener/internal/app/url_generator"
	"github.com/stretchr/testify/assert"
)

// func Test_ServeHTTP(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		method     string
// 		request    string
// 		body       string
// 		expectedCT string
// 		expectedSC int
// 		expectedRL string
// 	}{
// 		{
// 			name:       "test_1",
// 			method:     http.MethodPost,
// 			request:    "/",
// 			body:       "https://practicum.yandex.ru/",
// 			expectedCT: "application/json",
// 			expectedSC: http.StatusCreated,
// 			expectedRL: "",
// 		},
// 		{
// 			name:       "test_2_get_valid",
// 			method:     http.MethodGet,
// 			request:    "/aHR0cHM6",
// 			body:       "",
// 			expectedCT: "",
// 			expectedSC: http.StatusTemporaryRedirect,
// 			expectedRL: "https://practicum.yandex.ru/",
// 		},
// 		{
// 			name:       "test_3_get_invalid",
// 			method:     http.MethodGet,
// 			request:    "/invalidhash",
// 			body:       "",
// 			expectedCT: "text/plain; charset=utf-8",
// 			expectedSC: http.StatusBadRequest,
// 			expectedRL: "",
// 		},
// 	}
// 	generatorUrl := url_generator.Generator{Lenght: 8}
// 	storage := storage.Instance(generatorUrl)
// 	baseHandler := BaseHandler{
// 		BaseUrl: "http://localhost:8080",
// 		Storage: storage,
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {

// 			req := httptest.NewRequest(test.method, test.request, strings.NewReader(test.body))
// 			req.Header.Set("Content-Type", "text/plain")
// 			w := httptest.NewRecorder()

// 			baseHandler.ServeHTTP(w, req)

// 			resp := w.Result()

// 			defer resp.Body.Close()

// 			assert.Equal(t, test.expectedSC, resp.StatusCode)
// 			assert.Equal(t, test.expectedCT, resp.Header.Get("Content-Type"))

// 			if test.expectedRL != "" {
// 				assert.Equal(t, test.expectedRL, resp.Header.Get("Location"))
// 			}

// 			bodyBytes, err := io.ReadAll(resp.Body)
// 			assert.NoError(t, err)
// 			assert.Contains(t, string(bodyBytes), test.expectedRL)
// 		})
// 	}
// }

func Test_GetShortUrl(t *testing.T) {
	generatorUrl := url_generator.Generator{Lenght: 8}
	storage := storage.Instance(generatorUrl)
	baseHandler := BaseHandler{
		BaseUrl: "http://localhost:8080",
		Storage: storage,
	}

	t.Run("test_post_valid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://practicum.yandex.ru/"))
		req.Header.Set("Content-Type", "text/plain")
		w := httptest.NewRecorder()

		baseHandler.GetShortUrl(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		bodyBytes, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(bodyBytes), "http://localhost:8080/")
	})
}

func Test_GetFullUrl(t *testing.T) {
	generatorUrl := url_generator.Generator{Lenght: 8}
	storage := storage.Instance(generatorUrl)
	baseHandler := BaseHandler{
		BaseUrl: "http://localhost:8080",
		Storage: storage,
	}

	hash, _ := storage.SaveUrl("https://practicum.yandex.ru/")

	t.Run("test_get_valid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/"+hash, nil)
		w := httptest.NewRecorder()

		baseHandler.GetFullUrl(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode)
		assert.Equal(t, "https://practicum.yandex.ru/", resp.Header.Get("Location"))
	})

	t.Run("test_get_invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/invalidhash", nil)
		w := httptest.NewRecorder()

		baseHandler.GetFullUrl(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	})
}
