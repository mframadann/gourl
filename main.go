package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mframadann/gourl/configs"
	"github.com/mframadann/gourl/domain/link/controllers"
)

func main() {
	db := configs.InitDB()

	e := echo.New()
	rV_1 := e.Group("api/v1/")

	linkCtrler := controllers.NewItemController(db)

	rV_1.POST("shortlink/create", linkCtrler.Create)
	e.Start(":8080")
}
