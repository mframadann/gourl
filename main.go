package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/mframadann/gourl/configs"
	gropControllers "github.com/mframadann/gourl/domain/group-link/controllers"
	"github.com/mframadann/gourl/domain/link/controllers"
	"github.com/mframadann/gourl/helpers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := configs.InitDB()

	e := echo.New()
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = helpers.UseReadableErrMsg

	linkEndpoint := controllers.NewItemController(db)
	groupLinkEndpoint := gropControllers.NewGroupLinkController(db)

	rV_1 := e.Group("api/v1/")
	rV_1.GET("shortlink/get-links", linkEndpoint.GetAll)
	rV_1.POST("shortlink/create", linkEndpoint.Create)
	rV_1.PUT("shortlink/update", linkEndpoint.Update)
	rV_1.DELETE("shortlink/del", linkEndpoint.Delete)
	// Group Link Endpints
	rV_1.GET("group/get-groups", groupLinkEndpoint.GetAll)
	rV_1.POST("group/create", groupLinkEndpoint.Create)
	rV_1.PUT("group/update", groupLinkEndpoint.Update)
	rV_1.DELETE("group/delete", groupLinkEndpoint.Delete)

	e.Start(":8080")
}
