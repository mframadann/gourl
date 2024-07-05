package controllers

import (
	"net/http"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mframadann/gourl/domain/link/models"
	"github.com/mframadann/gourl/domain/link/services"
	"github.com/mframadann/gourl/helpers"
	"gorm.io/gorm"
)

type LinkController struct {
	itemService services.LinkService
	validate    vl.Validate
}

func (controller LinkController) Create(c echo.Context) error {
	type payload struct {
		UserID     *uint  `json:"user_id"`
		GroupID    *uint  `json:"group_id" `
		ShortedURL string `json:"shorted_url"`
		OriginURL  string `json:"origin_url" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return err
	}

	if payloadValidator.ShortedURL == "" {
		payloadValidator.ShortedURL = helpers.GenerateRandomShortenedLink()
	}

	result := controller.itemService.Create(
		models.Link{
			UserID:     payloadValidator.UserID,
			GroupID:    payloadValidator.GroupID,
			ShortedURL: payloadValidator.ShortedURL,
			OriginURL:  payloadValidator.OriginURL,
		},
	)

	return c.JSON(http.StatusOK, result)
}

func NewItemController(db *gorm.DB) LinkController {
	service := services.NewLinkService(db)
	controller := LinkController{
		itemService: service,
		validate:    *vl.New(),
	}

	return controller
}
