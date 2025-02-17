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

func Test_ServeHTTP(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		request    string
		body       string
		expectedCT string
		expectedSC int
		expectedRL string
	}{
		{
			name:       "test_1",
			method:     http.MethodPost,
			request:    "/",
			body:       "https://practicum.yandex.ru/",
			expectedCT: "application/json",
			expectedSC: http.StatusCreated,
			expectedRL: "",
		},
		{
			name:       "test_2_get_valid",
			method:     http.MethodGet,
			request:    "/aHR0cHM6",
			body:       "",
			expectedCT: "",
			expectedSC: http.StatusTemporaryRedirect,
			expectedRL: "https://practicum.yandex.ru/",
		},
		{
			name:       "test_3_get_invalid",
			method:     http.MethodGet,
			request:    "/invalidhash",
			body:       "",
			expectedCT: "text/plain; charset=utf-8",
			expectedSC: http.StatusBadRequest,
			expectedRL: "",
		},
	}
	generatorUrl := url_generator.Generator{Lenght: 8}
	storage := storage.Instance(generatorUrl)
	baseHandler := BaseHandler{
		BaseUrl: "http://localhost:8080",
		Storage: storage,
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			req := httptest.NewRequest(test.method, test.request, strings.NewReader(test.body))
			req.Header.Set("Content-Type", "text/plain")
			w := httptest.NewRecorder()

			baseHandler.ServeHTTP(w, req)

			resp := w.Result()

			defer resp.Body.Close()

			assert.Equal(t, test.expectedSC, resp.StatusCode)
			assert.Equal(t, test.expectedCT, resp.Header.Get("Content-Type"))

			if test.expectedRL != "" {
				assert.Equal(t, test.expectedRL, resp.Header.Get("Location"))
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Contains(t, string(bodyBytes), test.expectedRL)
		})
	}
}
