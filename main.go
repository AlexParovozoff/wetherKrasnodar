package main

import (
	_ "database/sql"
	"fmt"
	"net/http"
	"time"
	"math"

	_ "github.com/lib/pq"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/JonathanMH/goClacks/echo"
)

func main() {
	e := echo.New()
	e.Use(goClacks.Terrify)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	service.WetherRouter()
}
