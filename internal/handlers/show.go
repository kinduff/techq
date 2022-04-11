package handlers

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/kinduff/tech_qa/config/db"
	"github.com/kinduff/tech_qa/internal/models"
	resources "github.com/kinduff/tech_qa/resources"
)

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFS(resources.Templates, "templates/show.gohtml", "templates/layout.gohtml")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/q/")

	var question models.Question
	result := db.DB.First(&question, id)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	err = tpl.Execute(w, question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
