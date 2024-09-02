package models

import (
	"time"
)

type Link struct {
	ID         uint   `json:"link_id" gorm:"column:link_id"`
	UserID     *uint  `json:"user_id" gorm:"column:user_id"`
	GroupID    *uint  `json:"group_id" gorm:"column:group_id"`
	LinkTitle  string `json:"link_title" gorm:"column:link_title"`
	ShortedURL string `json:"shorted_url" gorm:"column:shorted_url;unique"`
	OriginURL  string `json:"origin_url" gorm:"column:origin_url"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type APILink struct {
	ID         uint   `json:"link_id" gorm:"column:link_id"`
	LinkTitle  string `json:"link_title"`
	ShortedURL string `json:"shorted_url"`
	OriginURL  string `json:"origin_url"`
}

func (Link) TableName() string {
	return "tb_links"
}
