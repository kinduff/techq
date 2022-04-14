package handlers

import (
	"net/http"
	"text/template"

	"github.com/kinduff/tech_qa/db"
	"github.com/kinduff/tech_qa/internal/models"
	resources "github.com/kinduff/tech_qa/resources"
)

// IndexHandler handles the root path.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" || r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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
