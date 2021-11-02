package main

import (
	"fmt"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/likevintw/echo/restful"
)

func main() {
	fmt.Println("hello world")
	e := echo.New()
	handler := restful.EchoHandler{e}
	fmt.Println(reflect.TypeOf(handler))
}
