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

	e.GET("/secret1", func(c echo.Context) error {
		out, err := exec.Command("cat", "/var/run/dockersec").Output()
		if err != nil {
			return c.String(http.StatusOK, "Could not execute exec!")
		}
		return c.String(http.StatusOK, "Secret is: "+string(out)+"\n")
	})

	e.GET("/secret2", func(c echo.Context) error {
		secret := os.Getenv("dockersec")
		return c.String(http.StatusOK, "Secret is: "+secret+"\n")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
