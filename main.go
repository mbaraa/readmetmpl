package main

import (
	"embed"
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

	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infof("dashboard's server started at port %s\n", "8081")
	log.Fatalln(log.ErrorLevel, http.ListenAndServe(":"+"8081", nil))
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	page := templates.NewIndex().Render(templates.IndexProps{})
	_, _ = io.Copy(w, page)
}
