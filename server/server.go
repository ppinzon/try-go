package server

import "github.com/labstack/echo/v4"

// Server bla
type Server struct {
	e    *echo.Echo
	conn Storage
}

// NewServer bla
func NewServer(_ *Config) *Server {
	e := echo.New()
	return &Server{
		e:    e,
		conn: NewInMemoryStorage(),
	}
}

//Run bla
func (s *Server) Run() error {
	s.initRoutes()
	return s.e.Start(":1323")
}

func (s *Server) initRoutes() {
	s.e.GET("/", HelloWorld)
	s.e.POST("/users", CreateUser)
	s.e.GET("/users/:id", GetUser)
	s.e.DELETE("/users/:id", DeleteUser)
	// s.e.GET("/users/list", ListUsers)
}
