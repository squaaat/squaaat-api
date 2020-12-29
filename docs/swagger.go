package docs

import (
	"io/ioutil"

	"github.com/swaggo/swag"
)

type s struct{}

func (s *s) ReadDoc() string {
	b, err := ioutil.ReadFile("./swagger.yml")
	if err != nil {
		panic(err.Error())
	}
	return string(b)
}

func init() {
	swag.Register(swag.Name, &s{})
}
