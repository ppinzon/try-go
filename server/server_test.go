package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// "try-go/server"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{"id":1,"name":"Jon","surname":"Snow","age":20}`
)

func TestCreateUser(t *testing.T) {
	// Setup
	s := NewServer(&Config{})

	// e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)
	// h := &User

	// Assertions
	if assert.NoError(t, s.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		body := rec.Body.Bytes()
		createdUser := User{}
		err := json.Unmarshal(body, &createdUser)

		assert.NoError(t, err)
		// assert.Equal(t, userJSON, rec.Body.String())
	}
}
