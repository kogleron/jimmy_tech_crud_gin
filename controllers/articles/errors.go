package articles

import (
	"net/http"

	"jimmy_tech_crud_gin/internal/errors"
)

var ErrNotFound = errors.New("not found", http.StatusNotFound)
