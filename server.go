package pony

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// ServerTarget instructs server on what to run.
type ServerTarget int

const (
	// SitePreviewServer instructs server to run site preview.
	SitePreviewServer ServerTarget = iota
	// SiteMapViewServer instructs server to run site map view.
	SiteMapViewServer
)

const siteMapTmpl = `
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Pony Site Map</title>
  </head>
	<body>
		<h1>Site Map</h1>
		<ul>
			{{range .Pages}}
				<li>{{.Path}} - {{.Name}}</li>
			{{end}}
		</ul>
  </body>
</html>`

// Server supports site preview and site map view.
type Server struct {
	Addr         string // server address for the site preview or site map (ip:port)
	MetadataFile string
	Target       ServerTarget
	TemplatesDir string
	pony         *Pony
	tmpl         *template.Template // site map template
}

// Start launches server.
func (s *Server) Start() error {
	if err := s.validate(); err != nil {
		return err
	}

	if err := s.init(); err != nil {
		return err
	}

	mux := http.NewServeMux()
	for route, handler := range s.routes() {
		mux.Handle(route, handler)
	}

	server := &http.Server{
		Addr:           s.Addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Pony is listening at http://%s...", s.Addr)
	return server.ListenAndServe()
}

func (s *Server) validate() error {
	errs := make([]string, 0)

	if _, err := os.Stat(s.MetadataFile); err != nil {
		errs = append(errs, errors.WithMessage(err, "metadata file read failed").Error())
	}

	if _, err := os.Stat(s.TemplatesDir); err != nil {
		errs = append(errs, errors.WithMessage(err, "templates directory read failed").Error())
	}

	if len(errs) == 0 {
		return nil
	}

	emsg := strings.Join(errs, "; ")
	return errors.New(emsg)
}

func (s *Server) init() error {
	opts := []Option{
		MetadataFile(s.MetadataFile),
		TemplatesDir(s.TemplatesDir),
	}
	p := NewPony(opts...)
	if errs := p.LoadAll(); errs != nil {
		log.Println(errs)
		return errors.New("failed to initialize pony")
	}

	s.pony = p

	if s.Target == SiteMapViewServer {
		tmpl, err := template.New("sitemap").Parse(siteMapTmpl)
		if err != nil {
			return err
		}
		s.tmpl = tmpl
	}

	return nil
}

func (s *Server) routes() map[string]http.HandlerFunc {
	if s.Target == SitePreviewServer {
		return s.previewRoutes()
	}

	if s.Target == SiteMapViewServer {
		return s.siteMapRoutes()
	}

	return nil
}

func (s *Server) previewRoutes() map[string]http.HandlerFunc {
	return nil
}

func (s *Server) siteMapRoutes() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/": siteMapFunc(s.pony.meta, s.tmpl),
	}
}

func siteMapFunc(meta *Metadata, tmpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, meta)
	})
}
