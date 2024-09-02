package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/mframadann/gourl/configs"
	authControllers "github.com/mframadann/gourl/domain/auth/controllers"
	gropControllers "github.com/mframadann/gourl/domain/group-link/controllers"
	"github.com/mframadann/gourl/domain/link/controllers"
	"github.com/mframadann/gourl/middlewares"
	"github.com/mframadann/gourl/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := configs.InitDB()

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = utils.UseReadableErrMsg
	e.Use(echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))

	linkEndpoint := controllers.NewItemController(db)
	groupLinkEndpoint := gropControllers.NewGroupLinkController(db)
	authEndpoints := authControllers.NewAuthController(db)

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
	// Auth endpoint
	rV_1.POST("register", authEndpoints.Register)
	rV_1.POST("sign-in", authEndpoints.SignIn)

	e.Start(":8080")
}
