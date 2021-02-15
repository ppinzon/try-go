package server

type Config struct {
	Port          int    // Port at which server runs
	User          string // User for basic auth
	Password      string // Password for basic auth
	EnableLogging bool   // EnableLogging enables logging of request/response
}
