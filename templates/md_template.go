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
	DocsUrl          string   `json:"docsUrl"`
	IssuesUrl        string   `json:"issuesUrl"`
	BuildBadge       string   `json:"buildBadge"`
	BuildStatusUrl   string   `json:"buildStatusUrl"`
	Description      string   `json:"description"`
	BuiltWith        []Anchor `json:"builtWith"`
	TechnologiesDocs []Anchor `json:"technologiesDocs"`
	PrerequisitesMd  string   `json:"prerequisites"`
	InstallationMd   string   `json:"installation"`
	RunningMd        string   `json:"running"`
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
