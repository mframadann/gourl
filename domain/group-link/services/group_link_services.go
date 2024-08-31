package services

import (
	"fmt"

	"github.com/mframadann/gourl/domain/group-link/models"
	"github.com/mframadann/gourl/domain/group-link/repositories"
	"github.com/mframadann/gourl/helpers"
	"gorm.io/gorm"
)

type GroupLinkService interface {
	CreateNewGroup(group models.GroupLink) helpers.Response
	GetAll() helpers.Response
	Update(groupId uint, group models.GroupLink) helpers.Response
	Delete(groupId uint) helpers.Response
}

type groupLinkService struct {
	groupLinkRepo repositories.GroupLinkRepository
}

func (service *groupLinkService) CreateNewGroup(group models.GroupLink) helpers.Response {
	var response helpers.Response
	if err := service.groupLinkRepo.Create(group); err != nil {
		response.Message = "Failed to create a new group " + err.Error()
		return response
	}

	response.Message = "Successfully created new group."
	return response
}

func (service *groupLinkService) Delete(groupId uint) helpers.Response {
	var response helpers.Response
	if err := service.groupLinkRepo.Delete(groupId); err != nil {
		response.Message = fmt.Sprintf("Error when trying to delete group: %s", err.Error())
		return response
	}

	response.Message = "Group hasbeen deleted."
	return response
}

func (service *groupLinkService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.groupLinkRepo.GetAll()
	if err != nil {
		response.Message = fmt.Sprintf("Failed to get group: %s", err.Error())
	}

	response.Data = data
	response.Message = "Success get all groups."

	return response
}

func (service *groupLinkService) Update(groupId uint, group models.GroupLink) helpers.Response {
	var response helpers.Response
	if err := service.groupLinkRepo.Update(groupId, group); err != nil {
		response.Message = fmt.Sprint("Failed to update group ", err.Error())
		return response
	}

	response.Message = "Group hasbeen updated"
	return response
}

func NewGroupLinkService(db *gorm.DB) GroupLinkService {
	return &groupLinkService{groupLinkRepo: repositories.NewGroupLinkRepository(db)}
}
