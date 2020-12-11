package internal

import (
	"html/template"
	"log"
	"net/http"
)

// Router defines settings to run static site preview.
type Router struct {
	// loaded metadata
	meta *Meta

	// loaded template
	tmpl *template.Template
}

// LoadTemplate loads builder template.
func (r *Router) LoadTemplate(tmpldir string) error {
	tmpl, err := template.ParseGlob(tmpldir + "/*.html")
	if err != nil {
		return err
	}

	r.tmpl = tmpl

	return nil
}

// LoadMeta loads pages metadata.
func (r *Router) LoadMeta(metafile string) error {
	meta, err := LoadMeta(metafile)
	if err != nil {
		return err
	}

	r.meta = meta

	return nil
}

// PreviewRoutes prepares route handlers required in preview.
func (r *Router) PreviewRoutes() map[string]http.HandlerFunc {
	routes := make(map[string]http.HandlerFunc)

	for _, page := range r.meta.Pages {
		routes[page.Path] = pageHandler(page, r.tmpl)
	}

	return routes
}

func pageHandler(page Page, tmpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if t := tmpl.Lookup(page.Template); t == nil {
			log.Printf("template with name %s not found\n", page.Template)
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}

		if err := tmpl.ExecuteTemplate(w, page.Template, page.Properties); err != nil {
			log.Println(err)
			return
		}
	})
}

// SiteMap returns site map route handler.
func (r *Router) SiteMap() http.HandlerFunc {
	return nil
}
