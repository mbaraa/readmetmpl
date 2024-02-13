package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"readmetmpl/log"
	"readmetmpl/templates"
)

var (
	//go:embed resources/*
	res embed.FS
)

func main() {
	http.HandleFunc("/", handleHomePage)

	http.HandleFunc("/api/generate-readme", handleGenerateReadme)

	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infof("dashboard's server started at port %s\n", "8081")
	log.Fatalln(log.ErrorLevel, http.ListenAndServe(":"+"8081", nil))
}

func handleGenerateReadme(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var reqBody templates.ReadmeProps
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v\n", reqBody)

	w.Header().Set("Content-Type", "text/plain")
	readme := templates.NewReadme().Render(reqBody)
	_, _ = io.Copy(w, readme)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	page := templates.NewIndex().Render(templates.IndexProps{})
	_, _ = io.Copy(w, page)
}
