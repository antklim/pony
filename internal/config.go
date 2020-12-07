package internal

// Server defines settings to run static site preview.
type Server struct {
	// TCP address for the server to listen on in the form "host:port".
	Addr string
}
