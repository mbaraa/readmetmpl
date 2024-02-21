package main

import (
	"embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"readmetmpl/log"
	"readmetmpl/templates"
)

var (
	//go:embed resources/*
	res embed.FS

	// Version contains the application's version number,
	// It's set by ldflags on build time.
	Version = ""
)

func main() {
	version := flag.Bool("version", false, "display version")
	help := flag.Bool("help", false, "display this screen")
	serverPort := flag.String("port", "", "web server's hosting port")
	filePath := flag.String("file", "", "path to the JSON file")
	flag.Parse()

	var err error

	switch {
	case *version:
		displayVersion()
	case len(*serverPort) != 0:
		err = startWebServer()
	case len(*filePath) != 0:
		err = generateReadmeFromJson(*filePath, func(tmpl templates.ReadmeProps) (io.Reader, error) {
			readme := templates.NewReadme().Render(tmpl)
			return readme, nil
		})
	case *help:
		fallthrough
	default:
		flag.PrintDefaults()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func displayVersion() {
	if Version == "" {
		Version = "(built from source)"
	}
	fmt.Printf("readmetmpl %s\n", Version)
	os.Exit(0)
}

func generateReadmeFromJson[T any](filePath string, renderFunc func(T) (io.Reader, error)) error {
	if filePath == "" {
		return errors.New("empty file path")
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	var jsonData T
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}
	readme, err := renderFunc(jsonData)
	if err != nil {
		return err
	}
	file, err := os.Create("Readme.md")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, readme)
	if err != nil {
		return err
	}

	return nil
}

func startWebServer() error {
	http.HandleFunc("/", handleHomePage)

	http.HandleFunc("/api/generate-readme", handleGenerateReadme)

	http.Handle("/resources/", http.FileServer(http.FS(res)))
	log.Infof("dashboard's server started at port %s\n", "8081")
	return http.ListenAndServe(":"+"8081", nil)
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
