package main

import (
	"embed"
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

var (
	//go:embed resources
	resources embed.FS

	//go:embed data/data.csv
	data embed.FS

	pages = map[string]string{
		"/":    "resources/index.gohtml",
		"/123": "resources/index.gohtml",
	}
)

type ViewData struct {
	Name  string
	Price int
}

func main() {
	csvFile, err := data.Open("data/data.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		fmt.Println(line)
	}

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
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	in := []ViewData{
		{
			Name:  "foo",
			Price: 100,
		},
		{
			Name:  "bar",
			Price: 200,
		},
	}
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]

	err = tpl.Execute(w, pick)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
