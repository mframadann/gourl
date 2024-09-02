package services

import (
	"net/http"

	"github.com/mframadann/gourl/domain/group-link/models"
	"github.com/mframadann/gourl/domain/group-link/repositories"
	"github.com/mframadann/gourl/utils"
	"gorm.io/gorm"
)

type GroupLinkService interface {
	CreateNewGroup(group models.GroupLink) utils.Response
	GetAll() utils.Response
	Update(groupId uint, group models.GroupLink) utils.Response
	Delete(groupId uint) utils.Response
}

type groupLinkService struct {
	groupLinkRepo repositories.GroupLinkRepository
}

func (service *groupLinkService) CreateNewGroup(group models.GroupLink) utils.Response {
	var response utils.Response
	if err := service.groupLinkRepo.Create(group); err != nil {
		response.Message = "Oops! failed to create new group link."
		response.StatusCode = http.StatusBadRequest
		return response
	}

	response.Message = "Successfully created new group."
	response.StatusCode = http.StatusOK
	return response
}

func (service *groupLinkService) Delete(groupId uint) utils.Response {
	var response utils.Response
	if err := service.groupLinkRepo.Delete(groupId); err != nil {
		response.Message = "Oops! failed to delete group link."
		response.StatusCode = http.StatusBadRequest
		return response
	}

	response.Message = "Group hasbeen deleted."
	response.StatusCode = http.StatusOK
	return response
}

func (service *groupLinkService) GetAll() utils.Response {
	var response utils.Response
	data, err := service.groupLinkRepo.GetAll()

	if err != nil {
		response.Message = "Oops! record not found or empty"
		response.StatusCode = http.StatusBadRequest
		return response
	}

	response.Data = data
	response.Message = "Success get all groups."
	response.StatusCode = http.StatusOK
	return response
}

func (service *groupLinkService) Update(groupId uint, group models.GroupLink) utils.Response {
	var response utils.Response
	if err := service.groupLinkRepo.Update(groupId, group); err != nil {
		response.Message = "Oops! failed to update group link."
		response.StatusCode = http.StatusBadRequest
		return response
	}

	response.Message = "Group hasbeen updated"
	response.StatusCode = http.StatusOK
	return response
}

func NewGroupLinkService(db *gorm.DB) GroupLinkService {
	return &groupLinkService{groupLinkRepo: repositories.NewGroupLinkRepository(db)}
}
