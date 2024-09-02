package repositories

import (
	"github.com/mframadann/gourl/domain/auth/dto"
	"github.com/mframadann/gourl/domain/auth/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user dto.Register) error
	SignIn(payload dto.SignIn) (models.User, error)
}

type db struct {
	Conn *gorm.DB
}

func (db *db) Register(user dto.Register) error {
	var newUser = models.User{
		FirstName:    user.LastName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
		Password:     user.Password,
	}
	return db.Conn.Create(&newUser).Error
}

func (db *db) SignIn(payload dto.SignIn) (models.User, error) {
	var data models.User
	result := db.Conn.Where("email_address", payload.EmailAddress).First(&data)

	return data, result.Error
}

func NewAuthRepository(Conn *gorm.DB) AuthRepository {
	return &db{Conn: Conn}
}
