package repositories

import (
	"github.com/mframadann/gourl/domain/link/dto"
	"github.com/mframadann/gourl/domain/link/models"
	"gorm.io/gorm"
)

type LinkRepository interface {
	Create(link models.Link) error
	Update(linkId uint, link models.Link) error
	Delete(linkId uint) error
	GetAll(queries dto.GetLinkQueries) ([]models.APILink, error)
}

type dbItem struct {
	Conn *gorm.DB
}

func (db *dbItem) Create(link models.Link) error {
	return db.Conn.Create(&link).Error
}

func (db *dbItem) GetAll(queries dto.GetLinkQueries) ([]models.APILink, error) {
	var data []models.APILink

	if queries.OrderByGrup {
		result := db.Conn.Model(&models.Link{}).Order("group_id").Find(&data)
		return data, result.Error
	}

	result := db.Conn.Model(&models.Link{}).Find(&data)

	return data, result.Error
}

func (db *dbItem) Update(linkId uint, item models.Link) error {
	return db.Conn.Where("link_id", linkId).Updates(item).Error
}

func (db *dbItem) Delete(linkId uint) error {
	return db.Conn.Delete(&models.Link{ID: linkId}).Error
}

func NewLinkRepository(Conn *gorm.DB) LinkRepository {
	return &dbItem{Conn: Conn}
}
