package repositories

import (
	"github.com/mframadann/gourl/domain/group-link/models"
	"gorm.io/gorm"
)

type GroupLinkRepository interface {
	Create(group models.GroupLink) error
	Update(groupId uint, group models.GroupLink) error
	Delete(groupId uint) error
	GetAll() ([]models.GroupLink, error)
}

type dbItem struct {
	Conn *gorm.DB
}

func (db *dbItem) Create(group models.GroupLink) error {
	return db.Conn.Create(&group).Error
}

func (db *dbItem) GetAll() ([]models.GroupLink, error) {
	var data []models.GroupLink
	result := db.Conn.Model(&models.GroupLink{}).Find(&data)
	return data, result.Error
}

func (db *dbItem) Update(groupId uint, group models.GroupLink) error {
	return db.Conn.Where("group_id", groupId).Updates(group).Error
}

func (db *dbItem) Delete(groupId uint) error {
	return db.Conn.Delete(&models.GroupLink{ID: groupId}).Error
}

func NewGroupLinkRepository(Conn *gorm.DB) GroupLinkRepository {
	return &dbItem{Conn: Conn}
}
