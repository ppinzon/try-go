package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// HelloWorld bla
func HelloWorld(c echo.Context) error {
	for _, user := range UserMap {
		fmt.Println(user.Name)
	}
	return c.String(http.StatusOK, "Hello, World!")
}

// TODO: IMPLEMENT PART A

// EchoTest is test
// func ListUsers(c echo.Context) error {
// 	return c.JSON(http.StatusOK, UserMap[])
// }

// CreateUser creates a brand new user
func CreateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	u.ID = NextUserID
	NextUserID++
	UserMap[u.ID] = u
	return c.JSON(http.StatusOK, u)
}

// GetUser displays a user
func GetUser(c echo.Context) error {
	id := c.Param("id")
	idint, _ := strconv.Atoi(id)

	user, ok := UserMap[idint]
	if ok {
		return c.JSON(http.StatusOK, user)
	}
	return c.String(http.StatusNotFound, "User not found")
}

// DeleteUser deletes users
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	idint, _ := strconv.Atoi(id)

	user, ok := UserMap[idint]
	if ok {
		delete(UserMap, idint)
		return c.JSON(http.StatusOK, user)
	}
	return c.String(http.StatusNotFound, "user not found")
}
