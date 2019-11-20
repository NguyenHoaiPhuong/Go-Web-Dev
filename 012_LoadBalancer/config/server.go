package config

// Server structure
type Server struct {
	Name        string
	Scheme      string
	Host        string
	Port        string
	Connections int
}

// URL returns server url
func (server Server) URL() string {
	return server.Scheme + "://" + server.Host + ":" + server.Port
}
