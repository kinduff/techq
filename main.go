package main

import (
	"embed"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/kinduff/tech_qa/db"
	"github.com/kinduff/tech_qa/models"
)

var (
	//go:embed resources
	resources embed.FS

	pages = map[string]string{
		"/": "resources/index.gohtml",
	}
)

type ViewData struct {
	Name  string
	Price int
}

func main() {
	db.ConnectDatabase()
	godotenv.Load()
	handleArgs()
	log.Println("Server started in port 3000")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	page, ok := pages[r.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tpl, err := template.ParseFS(resources, page)
	if err != nil {
		log.Printf("page %s not found in pages cache...", r.RequestURI)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var result models.Question
	db.DB.Raw("SELECT * FROM questions ORDER BY RANDOM() LIMIT 1;").Scan(&result)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	err = tpl.Execute(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			db.ExecuteSeed(db.DB, args[1:]...)
			os.Exit(0)
		case "create":
			db.CreateDB(db.DB)
			os.Exit(0)
		}
	}
}
