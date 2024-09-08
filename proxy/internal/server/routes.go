package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HelloWorldHandler)
	// e.GET("/account", func(c echo.Context) error {
	// 	resp, err := http.Get("http://localhost:8022" + c.Request().URL.Path)

	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, "User service unavaliable")
	// 	}
	// 	return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
	// })
	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
