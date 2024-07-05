package services

import (
	"fmt"

	"github.com/mframadann/gourl/domain/link/models"
	"github.com/mframadann/gourl/domain/link/repositories"
	"github.com/mframadann/gourl/helpers"
	"gorm.io/gorm"
)

var response helpers.Response

type LinkService interface {
	Create(item models.Link) helpers.Response
	GetById(idLink uint) helpers.Response
	GetAll() helpers.Response
	Update(idLink uint, item models.Link) helpers.Response
	Delete(idLink uint) helpers.Response
}

type linkService struct {
	linkRepo repositories.LinkRepository
}

func (service *linkService) Create(link models.Link) helpers.Response {
	if err := service.linkRepo.Create(link); err != nil {
		response.Message = "Failed to create a new item " + err.Error()
		return response
	}

	response.Message = "Successfully created new link."
	return response
}

func (service *linkService) Delete(linkId uint) helpers.Response {
	if err := service.linkRepo.Delete(linkId); err != nil {
		response.Message = fmt.Sprintf("Error when trying to delete link: %s", err.Error())
		return response
	}

	response.Message = "Link hasbeen deleted."
	return response
}

func (service *linkService) GetAll() helpers.Response {
	data, err := service.linkRepo.GetAll()
	if err != nil {
		response.Message = fmt.Sprintf("Failed to get links: %s", err.Error())
	}

	response.Data = data
	response.Message = "Success get all links."

	return response
}

func (service *linkService) GetById(linkId uint) helpers.Response {
	data, err := service.linkRepo.GetById(linkId)
	if err != nil {
		response.Message = fmt.Sprintf("Error when trying to get link: %s", err.Error())
	}

	response.Data = data
	response.Message = "Success get link"
	return response
}

func (service *linkService) Update(linkId uint, item models.Link) helpers.Response {
	if err := service.linkRepo.Update(linkId, item); err != nil {
		response.Message = fmt.Sprint("Failed to update link ", err.Error())
		return response
	}

	response.Message = "Link hasbeen updated"
	return response
}

func NewLinkService(db *gorm.DB) LinkService {
	return &linkService{linkRepo: repositories.NewLinkRepository(db)}
}
