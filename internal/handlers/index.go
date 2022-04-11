package handlers

import (
	"net/http"
	"text/template"

	"github.com/kinduff/tech_qa/config/db"
	"github.com/kinduff/tech_qa/internal/models"
	resources "github.com/kinduff/tech_qa/resources"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFS(resources.Templates, "templates/index.gohtml", "templates/layout.gohtml")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var question models.Question
	db.DB.Raw("SELECT * FROM questions ORDER BY RANDOM() LIMIT 1;").Scan(&question)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	err = tpl.Execute(w, question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
