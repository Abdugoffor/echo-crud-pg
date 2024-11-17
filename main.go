package main

import (
	"errors"
	"log"
	"os"

	"git.sriss.uz/shared/shared_service/logger"
	"git.sriss.uz/shared/shared_service/response"
	"github.com/labstack/echo/v4"
)

type (
	Bool bool
)

type String string

func (s String) Println() {
	log.Println(s)
}

type Ints []int

type Int int

type WW struct {
}

func (w WW) Write(p []byte) (n int, err error) {
	println(string(p))
	return 0, nil
}

func main() {
	f, _ := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	l := logger.NewWithWriter(f)

	l.Println("Hello world")

	e := echo.New()

	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			c.Response().Before(func() {
				println(c.Path())
				// println("before response")
			})

			c.Response().After(func() {
				// println("after response")
			})

			return next(c)
		}
	})

	e.GET("/a", func(c echo.Context) error {
		return response.HTTPError(errors.New("error")).BadRequest()
		// return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/", func(c echo.Context) error {
		return response.HTTPError(errors.New("error")).BadRequest()
		// return c.String(http.StatusOK, "Hello, World!")
	})
	e.Start(":8080")
}

// type Point struct {
// 	X float64
// 	Y float64
// 	S string
// }

// p := Point{}

// n, err := fmt.Sscanf(`(1,2.2)`, "(%f,%f)", &p.X, &p.Y)
// log.Println(n)
// log.Println(err)
// log.Println(p)
