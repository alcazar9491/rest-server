package main

import (
  // "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/alcazar9491/rest-server/routes"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  routes.Task(e)

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

