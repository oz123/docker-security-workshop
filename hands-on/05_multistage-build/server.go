package main

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Multistage build!\n")
	})

	e.GET("/hack", func(c echo.Context) error {
		out, err := exec.Command("cat", "/etc/passwd").Output()
		if err != nil {
			return c.String(http.StatusOK, "Could not execute exec!")
		}
		return c.String(http.StatusOK, string(out))
	})

	e.GET("/hack2", func(c echo.Context) error {
		cwd, err := os.Getwd()
		if err != nil {
			return c.String(http.StatusOK, "Could not execute getcwd!")
		}
		return c.String(http.StatusOK, "hacked "+cwd)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
