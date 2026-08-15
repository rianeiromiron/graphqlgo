package web

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func NewHandler() http.Handler {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "template execution failed", http.StatusInternalServerError)
		}
	})
}
