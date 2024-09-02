package services

import (
	"net/http"

	"github.com/mframadann/gourl/domain/link/dto"
	"github.com/mframadann/gourl/domain/link/models"
	"github.com/mframadann/gourl/domain/link/repositories"
	"github.com/mframadann/gourl/utils"
	"gorm.io/gorm"
)

type LinkService interface {
	Create(item models.Link) utils.Response
	GetAll(queries dto.GetLinkQueries) utils.Response
	Update(idLink uint, item models.Link) utils.Response
	Delete(idLink uint) utils.Response
}

type linkService struct {
	linkRepo repositories.LinkRepository
}

func (service *linkService) Create(link models.Link) utils.Response {
	var response utils.Response
	if err := service.linkRepo.Create(link); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Oops! failed create link."
		return response
	}

	response.Message = "Successfully created new link."
	response.StatusCode = http.StatusOK
	return response
}

func (service *linkService) Delete(linkId uint) utils.Response {
	var response utils.Response
	if err := service.linkRepo.Delete(linkId); err != nil {
		response.Message = "Oops! failed to delete link"
		response.StatusCode = http.StatusBadRequest
		return response
	}

	response.Message = "Link hasbeen deleted."
	response.StatusCode = http.StatusOK
	return response
}

func (service *linkService) GetAll(queries dto.GetLinkQueries) utils.Response {
	var response utils.Response
	data, err := service.linkRepo.GetAll(queries)

	if err != nil {
		response.Message = "Oops! record not found or empty"
		response.StatusCode = http.StatusNoContent
		return response
	}

	response.Data = data
	response.StatusCode = http.StatusOK
	response.Message = "Success get all links."

	return response
}

func (service *linkService) Update(linkId uint, item models.Link) utils.Response {
	var response utils.Response
	if err := service.linkRepo.Update(linkId, item); err != nil {
		response.Message = "Failed to update link "
		response.StatusCode = http.StatusBadRequest
		return response
	}

	response.Message = "Link hasbeen updated"
	return response
}

func NewLinkService(db *gorm.DB) LinkService {
	return &linkService{linkRepo: repositories.NewLinkRepository(db)}
}
