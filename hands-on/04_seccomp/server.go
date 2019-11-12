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
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/hack", func(c echo.Context) error {
		out, err := exec.Command("cat", "/etc/passwd").Output()
		if err != nil {
			return c.String(http.StatusOK, "Could not execute exec!\n")
		}
		return c.String(http.StatusOK, string(out)+"\n")
	})

	e.GET("/hack2", func(c echo.Context) error {
		cwd, err := os.Getwd()
		if err != nil {
			return c.String(http.StatusOK, "Could not execute getcwd!\n")
		}
		return c.String(http.StatusOK, "hacked "+cwd+"\n")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
