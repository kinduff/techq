package handlers

import (
	"net/http"
	"text/template"

	log "github.com/sirupsen/logrus"

	"github.com/kinduff/tech_qa/config/db"
	resources "github.com/kinduff/tech_qa/internal"
	"github.com/kinduff/tech_qa/internal/models"
)

var (
	question models.Question

	pages = map[string]string{
		"/": "templates/index.gohtml",
	}
)

func RoutesHandler(w http.ResponseWriter, r *http.Request) {
	page, ok := pages[r.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tpl, err := template.ParseFS(resources.Templates, page)
	if err != nil {
		log.Printf("page %s - %s not found in pages cache...", r.RequestURI, page)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db.DB.Raw("SELECT * FROM questions ORDER BY RANDOM() LIMIT 1;").Scan(&question)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	err = tpl.Execute(w, question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
