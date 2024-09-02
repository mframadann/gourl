package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mframadann/gourl/domain/link/dto"
	"github.com/mframadann/gourl/domain/link/models"
	"github.com/mframadann/gourl/domain/link/services"
	"github.com/mframadann/gourl/utils"
	"gorm.io/gorm"
)

type LinkController struct {
	LinkService services.LinkService
}

func (controller LinkController) GetAll(c echo.Context) error {
	q := new(dto.GetLinkQueries)

	if err := c.Bind(q); err != nil {
		return err
	}

	resp := controller.LinkService.GetAll(*q)
	return c.JSON(http.StatusOK, resp)
}

func (controller LinkController) Create(c echo.Context) error {
	l := new(dto.CreateLinkPayload)

	if err := c.Bind(l); err != nil {
		return err
	}

	if err := c.Validate(l); err != nil {
		return err
	}

	if l.ShortedURL == "" {
		l.ShortedURL = utils.GenerateRandomShortenedLink()
	}

	result := controller.LinkService.Create(
		models.Link{
			GroupID:    l.GroupID,
			LinkTitle:  l.Title,
			ShortedURL: l.ShortedURL,
			OriginURL:  l.OriginURL,
		},
	)

	return c.JSON(result.StatusCode, result)
}

func (controller LinkController) Update(c echo.Context) error {
	l := new(dto.UpdateLinkPayload)

	if err := c.Bind(l); err != nil {
		return err
	}

	if err := c.Validate(l); err != nil {
		return err
	}

	result := controller.LinkService.Update(
		l.ID,
		models.Link{
			GroupID:    l.GroupID,
			LinkTitle:  l.Title,
			ShortedURL: l.ShortedURL,
			OriginURL:  l.OriginURL,
		},
	)

	return c.JSON(http.StatusAccepted, result)
}

func (controller LinkController) Delete(c echo.Context) error {
	l := new(dto.DeleteLinkPalyload)

	if err := c.Bind(l); err != nil {
		return err
	}

	if err := c.Validate(l); err != nil {
		return err
	}

	result := controller.LinkService.Delete(l.ID)
	return c.JSON(http.StatusOK, result)
}

func NewItemController(db *gorm.DB) LinkController {
	service := services.NewLinkService(db)
	controller := LinkController{
		LinkService: service,
	}

	return controller
}
