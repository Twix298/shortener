package url_generator

import (
	"encoding/base64"
	"math/rand"
	"time"
)

type Generator struct {
	Lenght int
}

func (g *Generator) EncodeURL(url string) string {
	str := base64.StdEncoding.EncodeToString([]byte(url))
	if len(str) < g.Lenght {
		padding := make([]byte, g.Lenght-len(str))
		rand.Seed(time.Now().UnixNano())
		const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
		for i := range padding {
			padding[i] = chars[rand.Intn(len(chars))]
		}
		str += string(padding)
	}
	str = str[0:g.Lenght]
	return str
}
