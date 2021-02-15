package server

import "github.com/labstack/echo/v4"

type Server struct {
	e *echo.Echo
	conn Storage
}

func NewServer(_ *Config) *Server {
	e := echo.New()
	return &Server{
		e: e,
		conn: NewInMemoryStorage(),
	}
}

func (s *Server) Run() error {
	s.initRoutes()
	return s.e.Start(":1323")
}

func (s *Server) initRoutes() {
	s.e.GET("/", HelloWorld)
}