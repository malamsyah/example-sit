package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	RunServer()
}

func RunServer() {
	e := echo.New()

	host := os.Getenv("HOST")
	e.Use(middleware.Logger())
	e.GET("/example", func(c echo.Context) error {
		resp, err := http.Get(host + "/fact")
		if err != nil {
			return c.JSON(500, `{"message": "Error calling external service"}`)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		respBody := make(map[string]interface{})

		json.Unmarshal(body, &respBody)

		return c.JSON(200, respBody)
	})

	e.Logger.Fatal(e.Start(":9090"))
}
