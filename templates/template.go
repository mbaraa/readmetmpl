/*
Package tmplrndr
This package is responsible of rendenring templates in a sane(type safe) way.
*/
package templates

import (
	"bytes"
	"embed"
	"errors"
	"io"
	"text/template"
)

var (
	//go:embed html/*
	templates embed.FS

	templatesPaths = map[string][]string{
		"index": {"html/index.html", "html/_imports.html"},
	}

	_ Template[IndexProps] = &indexTemplate{}
)

// TemplateProps is a TYPED pages props, so that all pages get their props
// without any funny business when matching names and types.
type TemplateProps interface {
	IndexProps
}

// Template is an interface that represents a renderable html template.
type Template[T TemplateProps] interface {
	// Render accepts a generic prop type T,
	// and renders the templates with its props into the returned reader.
	Render(props T) io.Reader
}

func renderTemplate(name string, props any) (io.Reader, error) {
	var templatesPath []string
	if path, exists := templatesPaths[name]; !exists {
		return nil, errors.New("template doesn't exist")
	} else {
		templatesPath = path
	}
	tmpl := template.Must(template.ParseFS(templates, templatesPath...))
	out := bytes.NewBuffer([]byte{})
	err := tmpl.ExecuteTemplate(out, name, props)
	if err != nil {
		return nil, err
	}
	return out, nil
}
