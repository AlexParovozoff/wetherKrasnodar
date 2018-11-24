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

type (
	//current?lat={lat}&lon={lon}
	Wether struct {
		CityName    string    `json:"city_name"`
		Lat         float32   `json:"lat"`
		Lon         float32   `json:"lon"`
		Descriotion string    `json:"description"`
		ObjectTime  time.Time `json:"ob_time"`
		Temp
		Sun
		Wind
	}
	Temp struct {
		Temp    float32 `json:"temp"`
		TempApp float32 `json:"app_temp"`
	}
	Wind struct {
		WindCdir  string  `json:"wind_cdir"`
		WindSpeed float32 `json:"wind_speed"`
	}
	Sun struct {
		Sunrize string `json:"sunrise"`
		Sunset  string `json:"sunnset"`
	}
)

var(
	WetherMap = make([float64]Wether , 128 )
)

func WetherRouter() {
	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "Hellow user!")
	})
	e.POST("/", func(c echo.Context) error {
		return c.JSON(200, "Push coordinate")
	})
	e.GET("/:userid", func(c echo.Context) error {
		return c.JSON(200, ""+ userid)
	})
}