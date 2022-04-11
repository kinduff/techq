package handlers

import (
	"net/http"

	resources "github.com/kinduff/tech_qa/resources"
)

func StaticHandler() http.Handler {
	var staticFS = http.FS(resources.Statics)
	return http.FileServer(staticFS)
}
