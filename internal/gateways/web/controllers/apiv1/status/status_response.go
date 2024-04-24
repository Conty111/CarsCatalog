package status

import (
	"github.com/Conty111/CarsCatalog/internal/app/build"
)

// ResponseDoc is a response declaration for documentatino pruposes
type ResponseDoc struct {
	Data struct {
		Attributes Response `json:"attributes"`
	} `json:"data"`
}

// Response is a declaration for a status response
type Response struct {
	ID     string      `jsonapi:"primary,status"`
	Status string      `jsonapi:"attr,status"`
	Build  *build.Info `jsonapi:"attr,build"`
}
