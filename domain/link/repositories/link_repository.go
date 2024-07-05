package repositories

import (
	"github.com/mframadann/gourl/domain/link/models"
	"gorm.io/gorm"
)

type LinkRepository interface {
	Create(link models.Link) error
	Update(linkId uint, link models.Link) error
	Delete(linkId uint) error
	GetById(linkId uint) (models.Link, error)
	GetAll() ([]models.Link, error)
}

type dbItem struct {
	Conn *gorm.DB
}

func (db *dbItem) Create(link models.Link) error {
	return db.Conn.Create(&link).Error
}

func (db *dbItem) Delete(linkId uint) error {
	return db.Conn.Delete(&models.Link{ID: linkId}).Error
}

func (db *dbItem) GetAll() ([]models.Link, error) {
	var data []models.Link
	result := db.Conn.Find(&data)
	return data, result.Error
}

func (db *dbItem) GetById(linkId uint) (models.Link, error) {
	var data models.Link
	result := db.Conn.Where("id_item", linkId).First(&data)
	return data, result.Error
}

func (db *dbItem) Update(linkId uint, item models.Link) error {
	return db.Conn.Where("id_item", linkId).Updates(item).Error
}

func NewLinkRepository(Conn *gorm.DB) LinkRepository {
	return &dbItem{Conn: Conn}
}
