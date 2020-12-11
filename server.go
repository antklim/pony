package pony

// Server supports site preview and site map view.

type serverTarget int

const (
	sitePreview serverTarget = iota
	siteMapView
)

type server struct {
	addr   string // server address for the site preview or site map (ip:port)
	target serverTarget
	pony   Pony
}
