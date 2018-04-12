package main

import (
	"github.com/labstack/echo"
	control "github.com/wangxianzhuo/multi-control"
	"github.com/wangxianzhuo/multi-control/sample/print"
)

func main() {
	e := echo.New()
	m, err := control.NewMachine(make(chan string, 1))
	if err != nil {
		panic(err)
	}
	m.RegisterAsEchoMiddleware(e)
	u := print.Print{PID: "1"}
	m.AddUnit(&u)
	control.Routes(e)

	e.Start(":9999")
}
