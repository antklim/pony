package pony

import (
	"os"

	"github.com/pkg/errors"
)

// Server supports site preview and site map view.

type ServerTarget int

const (
	SitePreviewServer ServerTarget = iota
	SiteMapViewServer
)

type Server struct {
	Addr         string // server address for the site preview or site map (ip:port)
	MetadataFile string
	Target       ServerTarget
	TemplatesDir string
}

func (s *Server) Start() error {
	if _, err := os.Stat(s.MetadataFile); err != nil {
		return errors.Wrap(err, "metadata file read failed")
	}

	if _, err := os.Stat(s.TemplatesDir); err != nil {
		return errors.Wrap(err, "template file read failed")
	}

	// if err := router.LoadMeta(meta); err != nil {
	// 	return err
	// }

	// if err := router.LoadTemplate(tmpl); err != nil {
	// 	return err
	// }

	// mux := http.NewServeMux()

	// for route, handler := range router.PreviewRoutes() {
	// 	mux.Handle(route, handler)
	// }

	// s := &http.Server{
	// 	Addr:           addr,
	// 	Handler:        mux,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// log.Printf("pony preview is listening at %s", addr)
	// log.Fatal(s.ListenAndServe())
	return nil
}

// PreviewRoutes

// SiteMapRoutes
