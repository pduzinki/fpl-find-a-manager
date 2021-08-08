package html

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"net/http"
)

// Page is a helper type, holding pointer to template.Template
type Page struct {
	template *template.Template
}

//go:embed templates/*
var templates embed.FS

//go:embed static/*
var staticFiles embed.FS

// NewPage creates an html template based on files found in /static that match provided patterns
func NewPage(patterns ...string) (*Page, error) {
	patterns = append(patterns, "templates/layouts/*")
	t, err := template.ParseFS(templates, patterns...)
	if err != nil {
		return nil, err
	}

	return &Page{t}, nil
}

// Render executes page's html template and writes it to w
func (p *Page) Render(w http.ResponseWriter) {
	var buf bytes.Buffer
	err := p.template.ExecuteTemplate(w, "main.html", nil)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	io.Copy(w, &buf)
}

// StaticFiles returns http.FileServer containing static files (e.g. CSS style)
func StaticFiles() http.Handler {
	staticFS := http.FS(staticFiles)
	fs := http.FileServer(staticFS)

	return fs
}
