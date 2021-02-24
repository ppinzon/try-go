package server

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server bla
type Server struct {
	e          *echo.Echo
	conn       Storage
	NextUserID int
	UserMap    map[int]*User
}

// NewServer bla
func NewServer(_ *Config) *Server {
	e := echo.New()
	e.Use(middleware.Logger())

	// extUserID = 1

	return &Server{
		e:          e,
		conn:       NewInMemoryStorage(),
		UserMap:    make(map[int]*User),
		NextUserID: 1,
	}
}

//Run bla
func (s *Server) Run() error {
	s.auth()
	s.initRoutes()
	return s.e.Start(":1323")
}

func (s *Server) initRoutes() {
	s.e.GET("/", HelloWorld)
	s.e.POST("/users", s.CreateUser)
	s.e.GET("/users/:id", s.GetUser)
	s.e.DELETE("/users/:id", s.DeleteUser)
	// s.e.GET("/users/list", ListUsers)
}

func (s *Server) auth() {
	s.e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	}))
}
