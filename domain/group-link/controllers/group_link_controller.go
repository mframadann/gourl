package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mframadann/gourl/domain/group-link/dto"
	"github.com/mframadann/gourl/domain/group-link/models"
	"github.com/mframadann/gourl/domain/group-link/services"
	"gorm.io/gorm"
)

type GroupLinkController struct {
	GroupLinkService services.GroupLinkService
}

func (controller GroupLinkController) GetAll(c echo.Context) error {
	resp := controller.GroupLinkService.GetAll()
	return c.JSON(http.StatusOK, resp)
}

func (controller GroupLinkController) Create(c echo.Context) error {
	l := new(dto.CreateGroupPayload)

	if err := c.Bind(l); err != nil {
		return err
	}

	if err := c.Validate(l); err != nil {
		return err
	}

	result := controller.GroupLinkService.CreateNewGroup(
		models.GroupLink{
			GroupName: l.GroupName,
		},
	)

	return c.JSON(http.StatusOK, result)
}

func (controller GroupLinkController) Update(c echo.Context) error {
	l := new(dto.UpdateGroupPayload)

	if err := c.Bind(l); err != nil {
		return err
	}

	if err := c.Validate(l); err != nil {
		return err
	}

	result := controller.GroupLinkService.Update(
		l.ID,
		models.GroupLink{
			GroupName: l.GroupName,
		},
	)

	return c.JSON(http.StatusAccepted, result)
}

func (controller GroupLinkController) Delete(c echo.Context) error {
	l := new(dto.DeleteGroupPalyload)

	if err := c.Bind(l); err != nil {
		return err
	}

	if err := c.Validate(l); err != nil {
		return err
	}

	result := controller.GroupLinkService.Delete(l.ID)
	return c.JSON(http.StatusOK, result)
}

func NewGroupLinkController(db *gorm.DB) GroupLinkController {
	service := services.NewGroupLinkService(db)
	controller := GroupLinkController{
		GroupLinkService: service,
	}

	return controller
}
