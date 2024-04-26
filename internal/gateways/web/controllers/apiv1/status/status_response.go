package status

import (
	"github.com/Conty111/CarsCatalog/internal/app/build"
)

// Response is a declaration for a status response
type Response struct {
	Status string      `json:"status"`
	Build  *build.Info `json:"build"`
}
