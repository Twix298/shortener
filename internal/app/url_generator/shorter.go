package url_generator

import (
	"encoding/base64"
)

type Generator struct {
	Lenght int
}

func (g *Generator) EncodeURL(url string) string {
	str := base64.StdEncoding.EncodeToString([]byte(url))
	str = str[0:g.Lenght]
	return str
}
