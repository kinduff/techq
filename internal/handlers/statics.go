package handlers

import (
	"net/http"

	resources "github.com/kinduff/techq/resources"
)

// StaticHandler handles static files.
func StaticHandler() http.Handler {
	var staticFS = http.FS(resources.Statics)
	return http.FileServer(staticFS)
}
