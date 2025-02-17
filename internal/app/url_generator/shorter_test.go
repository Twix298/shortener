package url_generator

import (
	"testing"
)

func Test_EncodeURL(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "test_1",
			value: "www.yandex.ru",
			want:  "d3d3Lnlh",
		},
		{
			name:  "test_1",
			value: "www.google.com",
			want:  "d3d3Lmdv",
		},
	}

	for _, test := range tests {
		generator := Generator{Lenght: 8}
		t.Run(test.name, func(t *testing.T) {
			got := generator.EncodeURL(test.value)
			if got != test.want {
				t.Errorf("encodeUrl: url = %v, want %v", got, test.want)
			}

		})
	}
}
