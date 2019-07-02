package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    //"github.com/labstack/echo/middleware"

    //"./handler"
)

func main() {
    // Echo의 인스턴스를 만든다
    e := echo.New()

    // 모든 요청에 들어가는 미들웨어(로그 등)은 여기에
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // 라우팅
    //e.GET("/hello", handle.MainPage())

    e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
    e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
    e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

    // 서버 시작
    e.Run(standard.New(":1323"))
	//e.Run(standard.New(":8000"))
}