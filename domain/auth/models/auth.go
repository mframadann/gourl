package models

import "time"

type User struct {
	ID           uint      `json:"user_id" gorm:"column:user_id"`
	FirstName    string    `json:"first_name" gorm:"column:first_name"`
	LastName     string    `json:"last_name" gorm:"column:last_name"`
	EmailAddress string    `json:"email_address" gorm:"column:email_address"`
	Password     string    `json:"password" gorm:"column:password"`
	CreatedAt    time.Time `gorm:"column:registered_at;type:date"`
}

func (User) TableName() string {
	return "tb_users"
}
