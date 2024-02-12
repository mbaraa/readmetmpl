package templates

import (
	"bytes"
	"io"
	"readmetmpl/log"
)

type Anchor struct {
	Title string `json:"title"`
	Link  string `json:"url"`
}

type ReadmeProps struct {
	Title            string   `json:"title"`
	BriefDescription string   `json:"brief"`
	DocsUrl          string   `json:"docs-url"`
	IssuesUrl        string   `json:"issues-url"`
	BuildBadge       string   `json:"build-badge"`
	BuildStatusUrl   string   `json:"build-status-url"`
	Description      string   `json:"description"`
	BuiltWith        []Anchor `json:"built-with"`
	TechnologiesDocs []Anchor `json:"technologies-docs"`
	PrerequisitesMd  string   `json:"prerequisites-md"`
	InstallationMd   string   `json:"installation-md"`
	RunningMd        string   `json:"running-md"`
}

type readmeTemplate struct{}

// NewIndex returns a new index template instance.
func NewReadme() Template[ReadmeProps] {
	return &readmeTemplate{}
}

func (r *readmeTemplate) Render(props ReadmeProps) io.Reader {
	out, err := renderTextTemplate("readme", props)
	if err != nil {
		log.Errorln(err)
		return bytes.NewReader([]byte{})
	}
	return out
}
