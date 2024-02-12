/*
Package tmplrndr
This package is responsible of rendenring templates in a sane(type safe) way.
*/
package templates

import (
	"bytes"
	"embed"
	"errors"
	h "html/template"
	"io"
	t "text/template"
)

var (
	//go:embed html/*
	htmlTemplates embed.FS

	//go:embed md/*
	textTemplates embed.FS

	templatesPaths = map[string][]string{
		"index":  {"html/index.html", "html/_imports.html"},
		"readme": {"md/README.md"},
	}

	_ Template[IndexProps]  = &indexTemplate{}
	_ Template[ReadmeProps] = &readmeTemplate{}
)

// TemplateProps is a TYPED pages props, so that all pages get their props
// without any funny business when matching names and types.
type TemplateProps interface {
	IndexProps | ReadmeProps
}

// Template is an interface that represents a renderable html template.
type Template[T TemplateProps] interface {
	// Render accepts a generic prop type T,
	// and renders the templates with its props into the returned reader.
	Render(props T) io.Reader
}

func renderTextTemplate(name string, props any) (io.Reader, error) {
	return renderTemplate(name, props, md)
}

func renderHtmlTemplate(name string, props any) (io.Reader, error) {
	return renderTemplate(name, props, html)
}

type templateType string

const (
	html templateType = "html"
	md   templateType = "md"
)

func renderTemplate(name string, props any, renderer templateType) (io.Reader, error) {
	var templatesPath []string
	if path, exists := templatesPaths[name]; !exists {
		return nil, errors.New("template doesn't exist")
	} else {
		templatesPath = path
	}
	out := bytes.NewBuffer([]byte{})
	var err error
	switch renderer {
	case html:
		tmpl := h.Must(h.ParseFS(htmlTemplates, templatesPath...))
		err = tmpl.ExecuteTemplate(out, name, props)
	case md:
		tmpl := t.Must(t.ParseFS(textTemplates, templatesPath...))
		err = tmpl.ExecuteTemplate(out, name, props)
	}
	if err != nil {
		return nil, err
	}
	return out, nil
}
